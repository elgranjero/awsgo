package s3

// SelectObjectContent is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// This operation is not supported for directory buckets.
//
// This action filters the contents of an Amazon S3 object based on a simple
// structured query language (SQL) statement. In the request, along with the SQL
// expression, you must also specify a data serialization format (JSON, CSV, or
// Apache Parquet) of the object. Amazon S3 uses this format to parse object data
// into records, and returns only records that match the specified SQL expression.
// You must also specify the data serialization format for the response.
//
// This functionality is not supported for Amazon S3 on Outposts.
//
// For more information about Amazon S3 Select, see [Selecting Content from Objects] and [SELECT Command] in the Amazon S3 User
// Guide.
//
// Permissions You must have the s3:GetObject permission for this operation.
// Amazon S3 Select does not support anonymous access. For more information about
// permissions, see [Specifying Permissions in a Policy]in the Amazon S3 User Guide.
//
// Object Data Formats You can use Amazon S3 Select to query objects that have the
// following format properties:
//
// - CSV, JSON, and Parquet - Objects must be in CSV, JSON, or Parquet format.
//
// - UTF-8 - UTF-8 is the only encoding type Amazon S3 Select supports.
//
// - GZIP or BZIP2 - CSV and JSON files can be compressed using GZIP or BZIP2.
// GZIP and BZIP2 are the only compression formats that Amazon S3 Select supports
// for CSV and JSON files. Amazon S3 Select supports columnar compression for
// Parquet using GZIP or Snappy. Amazon S3 Select does not support whole-object
// compression for Parquet objects.
//
// - Server-side encryption - Amazon S3 Select supports querying objects that
// are protected with server-side encryption.
//
// For objects that are encrypted with customer-provided encryption keys (SSE-C),
//
// you must use HTTPS, and you must use the headers that are documented in the [GetObject].
// For more information about SSE-C, see [Server-Side Encryption (Using Customer-Provided Encryption Keys)]in the Amazon S3 User Guide.
//
// For objects that are encrypted with Amazon S3 managed keys (SSE-S3) and Amazon
//
// Web Services KMS keys (SSE-KMS), server-side encryption is handled
// transparently, so you don't need to specify anything. For more information about
// server-side encryption, including SSE-S3 and SSE-KMS, see [Protecting Data Using Server-Side Encryption]in the Amazon S3
// User Guide.
//
// Working with the Response Body Given the response size is unknown, Amazon S3
// Select streams the response as a series of messages and includes a
// Transfer-Encoding header with chunked as its value in the response. For more
// information, see [Appendix: SelectObjectContent Response].
//
// GetObject Support The SelectObjectContent action does not support the following
// GetObject functionality. For more information, see [GetObject].
//
// - Range : Although you can specify a scan range for an Amazon S3 Select
// request (see [SelectObjectContentRequest - ScanRange]in the request parameters), you cannot specify the range of
// bytes of an object to return.
//
// - The GLACIER , DEEP_ARCHIVE , and REDUCED_REDUNDANCY storage classes, or the
// ARCHIVE_ACCESS and DEEP_ARCHIVE_ACCESS access tiers of the INTELLIGENT_TIERING
// storage class: You cannot query objects in the GLACIER , DEEP_ARCHIVE , or
// REDUCED_REDUNDANCY storage classes, nor objects in the ARCHIVE_ACCESS or
// DEEP_ARCHIVE_ACCESS access tiers of the INTELLIGENT_TIERING storage class. For
// more information about storage classes, see [Using Amazon S3 storage classes]in the Amazon S3 User Guide.
//
// Special Errors For a list of special errors for this operation, see [List of SELECT Object Content Error Codes]
//
// The following operations are related to SelectObjectContent :
//
// [GetObject]
//
// [GetBucketLifecycleConfiguration]
//
// [PutBucketLifecycleConfiguration]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Appendix: SelectObjectContent Response]: https://docs.aws.amazon.com/AmazonS3/latest/API/RESTSelectObjectAppendix.html
// [Selecting Content from Objects]: https://docs.aws.amazon.com/AmazonS3/latest/dev/selecting-content-from-objects.html
// [PutBucketLifecycleConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketLifecycleConfiguration.html
// [SelectObjectContentRequest - ScanRange]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_SelectObjectContent.html#AmazonS3-SelectObjectContent-request-ScanRange
// [List of SELECT Object Content Error Codes]: https://docs.aws.amazon.com/AmazonS3/latest/API/ErrorResponses.html#SelectObjectContentErrorCodeList
// [GetBucketLifecycleConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketLifecycleConfiguration.html
// [Using Amazon S3 storage classes]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/storage-class-intro.html
// [SELECT Command]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-glacier-select-sql-reference-select.html
// [GetObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObject.html
// [Specifying Permissions in a Policy]: https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html
//
// [Server-Side Encryption (Using Customer-Provided Encryption Keys)]: https://docs.aws.amazon.com/AmazonS3/latest/dev/ServerSideEncryptionCustomerKeys.html
// [Protecting Data Using Server-Side Encryption]: https://docs.aws.amazon.com/AmazonS3/latest/dev/serv-side-encryption.html
