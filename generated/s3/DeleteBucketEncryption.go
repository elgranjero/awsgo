package s3

// DeleteBucketEncryption is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// This implementation of the DELETE action resets the default encryption for the
// bucket as server-side encryption with Amazon S3 managed keys (SSE-S3).
//
// - General purpose buckets - For information about the bucket default
// encryption feature, see [Amazon S3 Bucket Default Encryption]in the Amazon S3 User Guide.
//
// - Directory buckets - For directory buckets, there are only two supported
// options for server-side encryption: SSE-S3 and SSE-KMS. For information about
// the default encryption configuration in directory buckets, see [Setting default server-side encryption behavior for directory buckets].
//
// Permissions
//
// - General purpose bucket permissions - The s3:PutEncryptionConfiguration
// permission is required in a policy. The bucket owner has this permission by
// default. The bucket owner can grant this permission to others. For more
// information about permissions, see [Permissions Related to Bucket Operations]and [Managing Access Permissions to Your Amazon S3 Resources].
//
// - Directory bucket permissions - To grant access to this API operation, you
// must have the s3express:PutEncryptionConfiguration permission in an IAM
// identity-based policy instead of a bucket policy. Cross-account access to this
// API operation isn't supported. This operation can only be performed by the
// Amazon Web Services account that owns the resource. For more information about
// directory bucket policies and permissions, see [Amazon Web Services Identity and Access Management (IAM) for S3 Express One Zone]in the Amazon S3 User Guide.
//
// HTTP Host header syntax  Directory buckets - The HTTP Host header syntax is
// s3express-control.region-code.amazonaws.com .
//
// The following operations are related to DeleteBucketEncryption :
//
// [PutBucketEncryption]
//
// [GetBucketEncryption]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [GetBucketEncryption]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketEncryption.html
// [PutBucketEncryption]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketEncryption.html
// [Setting default server-side encryption behavior for directory buckets]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-bucket-encryption.html
// [Amazon S3 Bucket Default Encryption]: https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html
// [Managing Access Permissions to Your Amazon S3 Resources]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html
// [Permissions Related to Bucket Operations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources
// [Amazon Web Services Identity and Access Management (IAM) for S3 Express One Zone]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-security-iam.html
