package s3

// DeleteBucketInventoryConfiguration is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// This operation is not supported for directory buckets.
//
// Deletes an S3 Inventory configuration (identified by the inventory ID) from the
// bucket.
//
// To use this operation, you must have permissions to perform the
// s3:PutInventoryConfiguration action. The bucket owner has this permission by
// default. The bucket owner can grant this permission to others. For more
// information about permissions, see [Permissions Related to Bucket Subresource Operations]and [Managing Access Permissions to Your Amazon S3 Resources].
//
// For information about the Amazon S3 inventory feature, see [Amazon S3 Inventory].
//
// Operations related to DeleteBucketInventoryConfiguration include:
//
// [GetBucketInventoryConfiguration]
//
// [PutBucketInventoryConfiguration]
//
// [ListBucketInventoryConfigurations]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Amazon S3 Inventory]: https://docs.aws.amazon.com/AmazonS3/latest/dev/storage-inventory.html
// [ListBucketInventoryConfigurations]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListBucketInventoryConfigurations.html
// [Permissions Related to Bucket Subresource Operations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources
// [Managing Access Permissions to Your Amazon S3 Resources]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html
// [PutBucketInventoryConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketInventoryConfiguration.html
// [GetBucketInventoryConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketInventoryConfiguration.html
