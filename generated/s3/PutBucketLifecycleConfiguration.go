package s3

// PutBucketLifecycleConfiguration is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// Creates a new lifecycle configuration for the bucket or replaces an existing
// lifecycle configuration. Keep in mind that this will overwrite an existing
// lifecycle configuration, so if you want to retain any configuration details,
// they must be included in the new lifecycle configuration. For information about
// lifecycle configuration, see [Managing your storage lifecycle].
//
// Bucket lifecycle configuration now supports specifying a lifecycle rule using
// an object key name prefix, one or more object tags, object size, or any
// combination of these. Accordingly, this section describes the latest API. The
// previous version of the API supported filtering based only on an object key name
// prefix, which is supported for backward compatibility. For the related API
// description, see [PutBucketLifecycle].
//
// Rules You specify the lifecycle configuration in your request body. The
// lifecycle configuration is specified as XML consisting of one or more rules. An
// Amazon S3 Lifecycle configuration can have up to 1,000 rules. This limit is not
// adjustable.
//
// Bucket lifecycle configuration supports specifying a lifecycle rule using an
// object key name prefix, one or more object tags, object size, or any combination
// of these. Accordingly, this section describes the latest API. The previous
// version of the API supported filtering based only on an object key name prefix,
// which is supported for backward compatibility for general purpose buckets. For
// the related API description, see [PutBucketLifecycle].
//
// Lifecyle configurations for directory buckets only support expiring objects and
// cancelling multipart uploads. Expiring of versioned objects,transitions and tag
// filters are not supported.
//
// A lifecycle rule consists of the following:
//
// - A filter identifying a subset of objects to which the rule applies. The
// filter can be based on a key name prefix, object tags, object size, or any
// combination of these.
//
// - A status indicating whether the rule is in effect.
//
// - One or more lifecycle transition and expiration actions that you want
// Amazon S3 to perform on the objects identified by the filter. If the state of
// your bucket is versioning-enabled or versioning-suspended, you can have many
// versions of the same object (one current version and zero or more noncurrent
// versions). Amazon S3 provides predefined actions that you can specify for
// current and noncurrent object versions.
//
// For more information, see [Object Lifecycle Management] and [Lifecycle Configuration Elements].
//
// Permissions
// - General purpose bucket permissions - By default, all Amazon S3 resources
// are private, including buckets, objects, and related subresources (for example,
// lifecycle configuration and website configuration). Only the resource owner
// (that is, the Amazon Web Services account that created it) can access the
// resource. The resource owner can optionally grant access permissions to others
// by writing an access policy. For this operation, a user must have the
// s3:PutLifecycleConfiguration permission.
//
// You can also explicitly deny permissions. An explicit deny also supersedes any
//
// other permissions. If you want to block users or accounts from removing or
// deleting objects from your bucket, you must deny them permissions for the
// following actions:
//
// - s3:DeleteObject
//
// - s3:DeleteObjectVersion
//
// - s3:PutLifecycleConfiguration
//
// For more information about permissions, see [Managing Access Permissions to Your Amazon S3 Resources].
//
// - Directory bucket permissions - You must have the
// s3express:PutLifecycleConfiguration permission in an IAM identity-based policy
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
// The following operations are related to PutBucketLifecycleConfiguration :
//
// [GetBucketLifecycleConfiguration]
//
// [DeleteBucketLifecycle]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Object Lifecycle Management]: https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html
// [Lifecycle Configuration Elements]: https://docs.aws.amazon.com/AmazonS3/latest/dev/intro-lifecycle-rules.html
// [GetBucketLifecycleConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketLifecycleConfiguration.html
// [Authorizing Regional endpoint APIs with IAM]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-security-iam.html
// [PutBucketLifecycle]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketLifecycle.html
// [Managing Access Permissions to Your Amazon S3 Resources]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html
// [DeleteBucketLifecycle]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketLifecycle.html
// [Managing your storage lifecycle]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-lifecycle-mgmt.html
//
// [Concepts for directory buckets in Local Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-lzs-for-directory-buckets.html
// [Regional and Zonal endpoints for directory buckets in Availability Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/endpoint-directory-buckets-AZ.html
