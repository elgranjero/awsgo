package s3

// ListBucketInventoryConfigurations is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// This operation is not supported for directory buckets.
//
// Returns a list of S3 Inventory configurations for the bucket. You can have up
// to 1,000 inventory configurations per bucket.
//
// This action supports list pagination and does not return more than 100
// configurations at a time. Always check the IsTruncated element in the response.
// If there are no more configurations to list, IsTruncated is set to false. If
// there are more configurations to list, IsTruncated is set to true, and there is
// a value in NextContinuationToken . You use the NextContinuationToken value to
// continue the pagination of the list by passing the value in continuation-token
// in the request to GET the next page.
//
// To use this operation, you must have permissions to perform the
// s3:GetInventoryConfiguration action. The bucket owner has this permission by
// default. The bucket owner can grant this permission to others. For more
// information about permissions, see [Permissions Related to Bucket Subresource Operations]and [Managing Access Permissions to Your Amazon S3 Resources].
//
// For information about the Amazon S3 inventory feature, see [Amazon S3 Inventory]
//
// The following operations are related to ListBucketInventoryConfigurations :
//
// [GetBucketInventoryConfiguration]
//
// [DeleteBucketInventoryConfiguration]
//
// [PutBucketInventoryConfiguration]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Amazon S3 Inventory]: https://docs.aws.amazon.com/AmazonS3/latest/dev/storage-inventory.html
// [Permissions Related to Bucket Subresource Operations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources
// [DeleteBucketInventoryConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketInventoryConfiguration.html
// [Managing Access Permissions to Your Amazon S3 Resources]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html
// [PutBucketInventoryConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketInventoryConfiguration.html
// [GetBucketInventoryConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketInventoryConfiguration.html
