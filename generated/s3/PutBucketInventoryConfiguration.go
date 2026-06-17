package s3

// PutBucketInventoryConfiguration is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// This operation is not supported for directory buckets.
//
// This implementation of the PUT action adds an S3 Inventory configuration
// (identified by the inventory ID) to the bucket. You can have up to 1,000
// inventory configurations per bucket.
//
// Amazon S3 inventory generates inventories of the objects in the bucket on a
// daily or weekly basis, and the results are published to a flat file. The bucket
// that is inventoried is called the source bucket, and the bucket where the
// inventory flat file is stored is called the destination bucket. The destination
// bucket must be in the same Amazon Web Services Region as the source bucket.
//
// When you configure an inventory for a source bucket, you specify the
// destination bucket where you want the inventory to be stored, and whether to
// generate the inventory daily or weekly. You can also configure what object
// metadata to include and whether to inventory all object versions or only current
// versions. For more information, see [Amazon S3 Inventory]in the Amazon S3 User Guide.
//
// You must create a bucket policy on the destination bucket to grant permissions
// to Amazon S3 to write objects to the bucket in the defined location. For an
// example policy, see [Granting Permissions for Amazon S3 Inventory and Storage Class Analysis].
//
// Permissions To use this operation, you must have permission to perform the
// s3:PutInventoryConfiguration action. The bucket owner has this permission by
// default and can grant this permission to others.
//
// The s3:PutInventoryConfiguration permission allows a user to create an [S3 Inventory] report
// that includes all object metadata fields available and to specify the
// destination bucket to store the inventory. A user with read access to objects in
// the destination bucket can also access all object metadata fields that are
// available in the inventory report.
//
// To restrict access to an inventory report, see [Restricting access to an Amazon S3 Inventory report] in the Amazon S3 User Guide.
// For more information about the metadata fields available in S3 Inventory, see [Amazon S3 Inventory lists]
// in the Amazon S3 User Guide. For more information about permissions, see [Permissions related to bucket subresource operations]and [Identity and access management in Amazon S3]
// in the Amazon S3 User Guide.
//
// PutBucketInventoryConfiguration has the following special errors:
//
// HTTP 400 Bad Request Error  Code: InvalidArgument
//
// Cause: Invalid Argument
//
// HTTP 400 Bad Request Error  Code: TooManyConfigurations
//
// Cause: You are attempting to create a new configuration but have already
// reached the 1,000-configuration limit.
//
// HTTP 403 Forbidden Error  Cause: You are not the owner of the specified bucket,
// or you do not have the s3:PutInventoryConfiguration bucket permission to set
// the configuration on the bucket.
//
// The following operations are related to PutBucketInventoryConfiguration :
//
// [GetBucketInventoryConfiguration]
//
// [DeleteBucketInventoryConfiguration]
//
// [ListBucketInventoryConfigurations]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Granting Permissions for Amazon S3 Inventory and Storage Class Analysis]: https://docs.aws.amazon.com/AmazonS3/latest/dev/example-bucket-policies.html#example-bucket-policies-use-case-9
// [Amazon S3 Inventory]: https://docs.aws.amazon.com/AmazonS3/latest/dev/storage-inventory.html
// [ListBucketInventoryConfigurations]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListBucketInventoryConfigurations.html
// [S3 Inventory]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/storage-inventory.html
// [Permissions related to bucket subresource operations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources
// [DeleteBucketInventoryConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketInventoryConfiguration.html
// [Identity and access management in Amazon S3]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html
// [Restricting access to an Amazon S3 Inventory report]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/example-bucket-policies.html#example-bucket-policies-use-case-10
// [Amazon S3 Inventory lists]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/storage-inventory.html#storage-inventory-contents
// [GetBucketInventoryConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketInventoryConfiguration.html
