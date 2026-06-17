package s3

// DeleteBucketPolicy is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// Deletes the policy of a specified bucket.
//
// Directory buckets - For directory buckets, you must make requests for this API
// operation to the Regional endpoint. These endpoints support path-style requests
// in the format https://s3express-control.region-code.amazonaws.com/bucket-name .
// Virtual-hosted-style requests aren't supported. For more information about
// endpoints in Availability Zones, see [Regional and Zonal endpoints for directory buckets in Availability Zones]in the Amazon S3 User Guide. For more
// information about endpoints in Local Zones, see [Concepts for directory buckets in Local Zones]in the Amazon S3 User Guide.
//
// Permissions If you are using an identity other than the root user of the Amazon
// Web Services account that owns the bucket, the calling identity must both have
// the DeleteBucketPolicy permissions on the specified bucket and belong to the
// bucket owner's account in order to use this operation.
//
// If you don't have DeleteBucketPolicy permissions, Amazon S3 returns a 403
// Access Denied error. If you have the correct permissions, but you're not using
// an identity that belongs to the bucket owner's account, Amazon S3 returns a 405
// Method Not Allowed error.
//
// To ensure that bucket owners don't inadvertently lock themselves out of their
// own buckets, the root principal in a bucket owner's Amazon Web Services account
// can perform the GetBucketPolicy , PutBucketPolicy , and DeleteBucketPolicy API
// actions, even if their bucket policy explicitly denies the root principal's
// access. Bucket owner root principals can only be blocked from performing these
// API actions by VPC endpoint policies and Amazon Web Services Organizations
// policies.
//
// - General purpose bucket permissions - The s3:DeleteBucketPolicy permission is
// required in a policy. For more information about general purpose buckets bucket
// policies, see [Using Bucket Policies and User Policies]in the Amazon S3 User Guide.
//
// - Directory bucket permissions - To grant access to this API operation, you
// must have the s3express:DeleteBucketPolicy permission in an IAM identity-based
// policy instead of a bucket policy. Cross-account access to this API operation
// isn't supported. This operation can only be performed by the Amazon Web Services
// account that owns the resource. For more information about directory bucket
// policies and permissions, see [Amazon Web Services Identity and Access Management (IAM) for S3 Express One Zone]in the Amazon S3 User Guide.
//
// HTTP Host header syntax  Directory buckets - The HTTP Host header syntax is
// s3express-control.region-code.amazonaws.com .
//
// # The following operations are related to DeleteBucketPolicy
//
// [CreateBucket]
//
// [DeleteObject]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Concepts for directory buckets in Local Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-lzs-for-directory-buckets.html
// [DeleteObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteObject.html
// [Using Bucket Policies and User Policies]: https://docs.aws.amazon.com/AmazonS3/latest/dev/using-iam-policies.html
// [CreateBucket]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateBucket.html
// [Regional and Zonal endpoints for directory buckets in Availability Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/endpoint-directory-buckets-AZ.html
// [Amazon Web Services Identity and Access Management (IAM) for S3 Express One Zone]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-security-iam.html
