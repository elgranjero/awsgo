package s3control

// CreateStorageLensGroup is generated as a reference stub.
// Executable command wiring lives under cmd/s3control.go.
//
// Creates a new S3 Storage Lens group and associates it with the specified
//
// Amazon Web Services account ID. An S3 Storage Lens group is a custom grouping of
// objects based on prefix, suffix, object tags, object size, object age, or a
// combination of these filters. For each Storage Lens group that you’ve created,
// you can also optionally add Amazon Web Services resource tags. For more
// information about S3 Storage Lens groups, see [Working with S3 Storage Lens groups].
//
// To use this operation, you must have the permission to perform the
// s3:CreateStorageLensGroup action. If you’re trying to create a Storage Lens
// group with Amazon Web Services resource tags, you must also have permission to
// perform the s3:TagResource action. For more information about the required
// Storage Lens Groups permissions, see [Setting account permissions to use S3 Storage Lens groups].
//
// For information about Storage Lens groups errors, see [List of Amazon S3 Storage Lens error codes].
//
// [Setting account permissions to use S3 Storage Lens groups]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/storage_lens_iam_permissions.html#storage_lens_groups_permissions
// [List of Amazon S3 Storage Lens error codes]: https://docs.aws.amazon.com/AmazonS3/latest/API/ErrorResponses.html#S3LensErrorCodeList
// [Working with S3 Storage Lens groups]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/storage-lens-groups-overview.html
