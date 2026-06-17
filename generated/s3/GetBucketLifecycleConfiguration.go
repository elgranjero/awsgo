package s3

// GetBucketLifecycleConfiguration is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// Returns the lifecycle configuration information set on the bucket. For
// information about lifecycle configuration, see [Object Lifecycle Management].
//
// Bucket lifecycle configuration now supports specifying a lifecycle rule using
// an object key name prefix, one or more object tags, object size, or any
// combination of these. Accordingly, this section describes the latest API, which
// is compatible with the new functionality. The previous version of the API
// supported filtering based only on an object key name prefix, which is supported
// for general purpose buckets for backward compatibility. For the related API
// description, see [GetBucketLifecycle].
//
// Lifecyle configurations for directory buckets only support expiring objects and
// cancelling multipart uploads. Expiring of versioned objects, transitions and tag
// filters are not supported.
//
// Permissions
// - General purpose bucket permissions - By default, all Amazon S3 resources
// are private, including buckets, objects, and related subresources (for example,
// lifecycle configuration and website configuration). Only the resource owner
// (that is, the Amazon Web Services account that created it) can access the
// resource. The resource owner can optionally grant access permissions to others
// by writing an access policy. For this operation, a user must have the
// s3:GetLifecycleConfiguration permission.
//
// For more information about permissions, see [Managing Access Permissions to Your Amazon S3 Resources].
//
// - Directory bucket permissions - You must have the
// s3express:GetLifecycleConfiguration permission in an IAM identity-based policy
// to use this operation. Cross-account access to this API operation isn't
// supported. The resource owner can optionally grant access permissions to others
// by creating a role or user for them as long as they are within the same account
// as the owner and resource.
//
// For more information about directory bucket policies and permissions, see [Authorizing Regional endpoint APIs with IAM]in
//
// the Amazon S3 User Guide.
//
// Directory buckets - For directory buckets, you must make requests for this API
//
// operation to the Regional endpoint. These endpoints support path-style requests
// in the format https://s3express-control.region-code.amazonaws.com/bucket-name
// . Virtual-hosted-style requests aren't supported. For more information about
// endpoints in Availability Zones, see [Regional and Zonal endpoints for directory buckets in Availability Zones]in the Amazon S3 User Guide. For more
// information about endpoints in Local Zones, see [Concepts for directory buckets in Local Zones]in the Amazon S3 User Guide.
//
// HTTP Host header syntax  Directory buckets - The HTTP Host header syntax is
// s3express-control.region.amazonaws.com .
//
// GetBucketLifecycleConfiguration has the following special error:
//
// - Error code: NoSuchLifecycleConfiguration
//
// - Description: The lifecycle configuration does not exist.
//
// - HTTP Status Code: 404 Not Found
//
// - SOAP Fault Code Prefix: Client
//
// The following operations are related to GetBucketLifecycleConfiguration :
//
// [GetBucketLifecycle]
//
// [PutBucketLifecycle]
//
// [DeleteBucketLifecycle]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [GetBucketLifecycle]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketLifecycle.html
// [Object Lifecycle Management]: https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html
// [Authorizing Regional endpoint APIs with IAM]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-security-iam.html
// [PutBucketLifecycle]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketLifecycle.html
// [Managing Access Permissions to Your Amazon S3 Resources]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html
// [DeleteBucketLifecycle]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketLifecycle.html
//
// [Concepts for directory buckets in Local Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-lzs-for-directory-buckets.html
// [Regional and Zonal endpoints for directory buckets in Availability Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/endpoint-directory-buckets-AZ.html
