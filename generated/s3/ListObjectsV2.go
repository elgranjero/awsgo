package s3

// ListObjectsV2 is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// Returns some or all (up to 1,000) of the objects in a bucket with each request.
// You can use the request parameters as selection criteria to return a subset of
// the objects in a bucket. A 200 OK response can contain valid or invalid XML.
// Make sure to design your application to parse the contents of the response and
// handle it appropriately. For more information about listing objects, see [Listing object keys programmatically]in the
// Amazon S3 User Guide. To get a list of your buckets, see [ListBuckets].
//
// - General purpose bucket - For general purpose buckets, ListObjectsV2 doesn't
// return prefixes that are related only to in-progress multipart uploads.
//
// - Directory buckets - For directory buckets, ListObjectsV2 response includes
// the prefixes that are related only to in-progress multipart uploads.
//
// - Directory buckets - For directory buckets, you must make requests for this
// API operation to the Zonal endpoint. These endpoints support
// virtual-hosted-style requests in the format
// https://amzn-s3-demo-bucket.s3express-zone-id.region-code.amazonaws.com/key-name
// . Path-style requests are not supported. For more information about endpoints
// in Availability Zones, see [Regional and Zonal endpoints for directory buckets in Availability Zones]in the Amazon S3 User Guide. For more information
// about endpoints in Local Zones, see [Concepts for directory buckets in Local Zones]in the Amazon S3 User Guide.
//
// Permissions
//
// - General purpose bucket permissions - To use this operation, you must have
// READ access to the bucket. You must have permission to perform the
// s3:ListBucket action. The bucket owner has this permission by default and can
// grant this permission to others. For more information about permissions, see [Permissions Related to Bucket Subresource Operations]
// and [Managing Access Permissions to Your Amazon S3 Resources]in the Amazon S3 User Guide.
//
// - Directory bucket permissions - To grant access to this API operation on a
// directory bucket, we recommend that you use the [CreateSession]CreateSession API operation
// for session-based authorization. Specifically, you grant the
// s3express:CreateSession permission to the directory bucket in a bucket policy
// or an IAM identity-based policy. Then, you make the CreateSession API call on
// the bucket to obtain a session token. With the session token in your request
// header, you can make API requests to this operation. After the session token
// expires, you make another CreateSession API call to generate a new session
// token for use. Amazon Web Services CLI or SDKs create session and refresh the
// session token automatically to avoid service interruptions when a session
// expires. For more information about authorization, see [CreateSession]CreateSession .
//
// Sorting order of returned objects
//
// - General purpose bucket - For general purpose buckets, ListObjectsV2 returns
// objects in lexicographical order based on their key names.
//
// - Directory bucket - For directory buckets, ListObjectsV2 does not return
// objects in lexicographical order.
//
// HTTP Host header syntax  Directory buckets - The HTTP Host header syntax is
// Bucket-name.s3express-zone-id.region-code.amazonaws.com .
//
// This section describes the latest revision of this action. We recommend that
// you use this revised API operation for application development. For backward
// compatibility, Amazon S3 continues to support the prior version of this API
// operation, [ListObjects].
//
// The following operations are related to ListObjectsV2 :
//
// [GetObject]
//
// [PutObject]
//
// [CreateBucket]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [ListObjects]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListObjects.html
// [Concepts for directory buckets in Local Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-lzs-for-directory-buckets.html
// [Permissions Related to Bucket Subresource Operations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources
// [Listing object keys programmatically]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/ListingKeysUsingAPIs.html
// [ListBuckets]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListBuckets.html
// [PutObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObject.html
// [Managing Access Permissions to Your Amazon S3 Resources]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html
// [CreateSession]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateSession.html
// [GetObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObject.html
// [CreateBucket]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateBucket.html
// [Regional and Zonal endpoints for directory buckets in Availability Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/endpoint-directory-buckets-AZ.html
