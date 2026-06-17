package s3

// DeleteBucket is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// Deletes the S3 bucket. All objects (including all object versions and delete
// markers) in the bucket must be deleted before the bucket itself can be deleted.
//
// - Directory buckets - If multipart uploads in a directory bucket are in
// progress, you can't delete the bucket until all the in-progress multipart
// uploads are aborted or completed.
//
// - Directory buckets - For directory buckets, you must make requests for this
// API operation to the Regional endpoint. These endpoints support path-style
// requests in the format
// https://s3express-control.region-code.amazonaws.com/bucket-name .
// Virtual-hosted-style requests aren't supported. For more information about
// endpoints in Availability Zones, see [Regional and Zonal endpoints for directory buckets in Availability Zones]in the Amazon S3 User Guide. For more
// information about endpoints in Local Zones, see [Concepts for directory buckets in Local Zones]in the Amazon S3 User Guide.
//
// Permissions
//
// - General purpose bucket permissions - You must have the s3:DeleteBucket
// permission on the specified bucket in a policy.
//
// - Directory bucket permissions - You must have the s3express:DeleteBucket
// permission in an IAM identity-based policy instead of a bucket policy.
// Cross-account access to this API operation isn't supported. This operation can
// only be performed by the Amazon Web Services account that owns the resource. For
// more information about directory bucket policies and permissions, see [Amazon Web Services Identity and Access Management (IAM) for S3 Express One Zone]in the
// Amazon S3 User Guide.
//
// HTTP Host header syntax  Directory buckets - The HTTP Host header syntax is
// s3express-control.region-code.amazonaws.com .
//
// The following operations are related to DeleteBucket :
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
// [CreateBucket]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateBucket.html
// [Regional and Zonal endpoints for directory buckets in Availability Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/endpoint-directory-buckets-AZ.html
// [Amazon Web Services Identity and Access Management (IAM) for S3 Express One Zone]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-security-iam.html
