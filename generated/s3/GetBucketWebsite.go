package s3

// GetBucketWebsite is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// This operation is not supported for directory buckets.
//
// Returns the website configuration for a bucket. To host website on Amazon S3,
// you can configure a bucket as website by adding a website configuration. For
// more information about hosting websites, see [Hosting Websites on Amazon S3].
//
// This GET action requires the S3:GetBucketWebsite permission. By default, only
// the bucket owner can read the bucket website configuration. However, bucket
// owners can allow other users to read the website configuration by writing a
// bucket policy granting them the S3:GetBucketWebsite permission.
//
// The following operations are related to GetBucketWebsite :
//
// [DeleteBucketWebsite]
//
// [PutBucketWebsite]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [PutBucketWebsite]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketWebsite.html
// [Hosting Websites on Amazon S3]: https://docs.aws.amazon.com/AmazonS3/latest/dev/WebsiteHosting.html
// [DeleteBucketWebsite]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketWebsite.html
