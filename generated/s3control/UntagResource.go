package s3control

// UntagResource is generated as a reference stub.
// Executable command wiring lives under cmd/s3control.go.
//
// This operation removes the specified user-defined tags from an S3 resource. You
// can pass one or more tag keys.
//
// This operation is only supported for the following Amazon S3 resources:
//
// [General purpose buckets]
//
// [Access Points for directory buckets]
//
// [Access Points for general purpose buckets]
//
// [Directory buckets]
//
// [S3 Storage Lens groups]
//
// [S3 Access Grants instances, registered locations, and grants]
// - .
//
// Permissions For general purpose buckets, access points for general purpose
// buckets, Storage Lens groups, and S3 Access Grants, you must have the
// s3:UntagResource permission to use this operation.
//
// Directory bucket permissions For directory buckets, you must have the
// s3express:UntagResource permission to use this operation. For more information
// about directory buckets policies and permissions, see [Identity and Access Management (IAM) for S3 Express One Zone]in the Amazon S3 User
// Guide.
//
// HTTP Host header syntax  Directory buckets - The HTTP Host header syntax is
// s3express-control.region.amazonaws.com .
//
// For information about S3 Tagging errors, see [List of Amazon S3 Tagging error codes].
//
// [Access Points for directory buckets]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/access-points-db-tagging.html
// [General purpose buckets]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/buckets-tagging.html
// [S3 Access Grants instances, registered locations, and grants]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/access-grants-tagging.html
// [List of Amazon S3 Tagging error codes]: https://docs.aws.amazon.com/AmazonS3/latest/API/ErrorResponses.html#S3TaggingErrorCodeList
// [Directory buckets]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/directory-buckets-tagging.html
// [Identity and Access Management (IAM) for S3 Express One Zone]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-permissions.html
// [Access Points for general purpose buckets]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/access-points-tagging.html
// [S3 Storage Lens groups]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/storage-lens-groups.html
