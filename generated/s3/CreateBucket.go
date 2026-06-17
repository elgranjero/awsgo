package s3

// CreateBucket is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// This action creates an Amazon S3 bucket. To create an Amazon S3 on Outposts
// bucket, see [CreateBucket]CreateBucket .
//
// Creates a new S3 bucket. To create a bucket, you must set up Amazon S3 and have
// a valid Amazon Web Services Access Key ID to authenticate requests. Anonymous
// requests are never allowed to create buckets. By creating the bucket, you become
// the bucket owner.
//
// There are two types of buckets: general purpose buckets and directory buckets.
// For more information about these bucket types, see [Creating, configuring, and working with Amazon S3 buckets]in the Amazon S3 User Guide.
//
// - General purpose buckets - If you send your CreateBucket request to the
// s3.amazonaws.com global endpoint, the request goes to the us-east-1 Region. So
// the signature calculations in Signature Version 4 must use us-east-1 as the
// Region, even if the location constraint in the request specifies another Region
// where the bucket is to be created. If you create a bucket in a Region other than
// US East (N. Virginia), your application must be able to handle 307 redirect. For
// more information, see [Virtual hosting of buckets]in the Amazon S3 User Guide.
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
// - General purpose bucket permissions - In addition to the s3:CreateBucket
// permission, the following permissions are required in a policy when your
// CreateBucket request includes specific headers:
//
// - Access control lists (ACLs) - In your CreateBucket request, if you specify
// an access control list (ACL) and set it to public-read , public-read-write ,
// authenticated-read , or if you explicitly specify any other custom ACLs, both
// s3:CreateBucket and s3:PutBucketAcl permissions are required. In your
// CreateBucket request, if you set the ACL to private , or if you don't specify
// any ACLs, only the s3:CreateBucket permission is required.
//
// - Object Lock - In your CreateBucket request, if you set
// x-amz-bucket-object-lock-enabled to true, the
// s3:PutBucketObjectLockConfiguration and s3:PutBucketVersioning permissions are
// required.
//
// - S3 Object Ownership - If your CreateBucket request includes the
// x-amz-object-ownership header, then the s3:PutBucketOwnershipControls
// permission is required.
//
// To set an ACL on a bucket as part of a CreateBucket request, you must explicitly
//
// set S3 Object Ownership for the bucket to a different value than the default,
// BucketOwnerEnforced . Additionally, if your desired bucket ACL grants public
// access, you must first create the bucket (without the bucket ACL) and then
// explicitly disable Block Public Access on the bucket before using PutBucketAcl
// to set the ACL. If you try to create a bucket with a public ACL, the request
// will fail.
//
// For the majority of modern use cases in S3, we recommend that you keep all
//
// Block Public Access settings enabled and keep ACLs disabled. If you would like
// to share data with users outside of your account, you can use bucket policies as
// needed. For more information, see [Controlling ownership of objects and disabling ACLs for your bucket]and [Blocking public access to your Amazon S3 storage]in the Amazon S3 User Guide.
//
// - S3 Block Public Access - If your specific use case requires granting public
// access to your S3 resources, you can disable Block Public Access. Specifically,
// you can create a new bucket with Block Public Access enabled, then separately
// call the [DeletePublicAccessBlock]DeletePublicAccessBlock API. To use this operation, you must have the
// s3:PutBucketPublicAccessBlock permission. For more information about S3 Block
// Public Access, see [Blocking public access to your Amazon S3 storage]in the Amazon S3 User Guide.
//
// - Directory bucket permissions - You must have the s3express:CreateBucket
// permission in an IAM identity-based policy instead of a bucket policy.
// Cross-account access to this API operation isn't supported. This operation can
// only be performed by the Amazon Web Services account that owns the resource. For
// more information about directory bucket policies and permissions, see [Amazon Web Services Identity and Access Management (IAM) for S3 Express One Zone]in the
// Amazon S3 User Guide.
//
// The permissions for ACLs, Object Lock, S3 Object Ownership, and S3 Block Public
//
// Access are not supported for directory buckets. For directory buckets, all Block
// Public Access settings are enabled at the bucket level and S3 Object Ownership
// is set to Bucket owner enforced (ACLs disabled). These settings can't be
// modified.
//
// For more information about permissions for creating and working with directory
//
// buckets, see [Directory buckets]in the Amazon S3 User Guide. For more information about
// supported S3 features for directory buckets, see [Features of S3 Express One Zone]in the Amazon S3 User Guide.
//
// HTTP Host header syntax  Directory buckets - The HTTP Host header syntax is
// s3express-control.region-code.amazonaws.com .
//
// The following operations are related to CreateBucket :
//
// [PutObject]
//
// [DeleteBucket]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Creating, configuring, and working with Amazon S3 buckets]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/creating-buckets-s3.html
// [Concepts for directory buckets in Local Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-lzs-for-directory-buckets.html
// [PutObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObject.html
// [DeleteBucket]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucket.html
// [CreateBucket]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateBucket.html
// [Virtual hosting of buckets]: https://docs.aws.amazon.com/AmazonS3/latest/dev/VirtualHosting.html
// [Regional and Zonal endpoints for directory buckets in Availability Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/endpoint-directory-buckets-AZ.html
//
// [DeletePublicAccessBlock]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeletePublicAccessBlock.html
// [Directory buckets]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/directory-buckets-overview.html
// [Features of S3 Express One Zone]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-one-zone.html#s3-express-features
// [Amazon Web Services Identity and Access Management (IAM) for S3 Express One Zone]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-security-iam.html
// [Controlling ownership of objects and disabling ACLs for your bucket]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/about-object-ownership.html
// [Blocking public access to your Amazon S3 storage]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/access-control-block-public-access.html
