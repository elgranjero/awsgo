package s3

// GetObjectAttributes is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// Retrieves all of the metadata from an object without returning the object
// itself. This operation is useful if you're interested only in an object's
// metadata.
//
// GetObjectAttributes combines the functionality of HeadObject and ListParts . All
// of the data returned with both of those individual calls can be returned with a
// single call to GetObjectAttributes .
//
// Directory buckets - For directory buckets, you must make requests for this API
// operation to the Zonal endpoint. These endpoints support virtual-hosted-style
// requests in the format
// https://amzn-s3-demo-bucket.s3express-zone-id.region-code.amazonaws.com/key-name
// . Path-style requests are not supported. For more information about endpoints
// in Availability Zones, see [Regional and Zonal endpoints for directory buckets in Availability Zones]in the Amazon S3 User Guide. For more information
// about endpoints in Local Zones, see [Concepts for directory buckets in Local Zones]in the Amazon S3 User Guide.
//
// Permissions
// - General purpose bucket permissions - To use GetObjectAttributes , you must
// have READ access to the object.
//
// The other permissions that you need to use this operation depend on whether the
//
// bucket is versioned and if a version ID is passed in the GetObjectAttributes
// request.
//
// - If you pass a version ID in your request, you need both the
// s3:GetObjectVersion and s3:GetObjectVersionAttributes permissions.
//
// - If you do not pass a version ID in your request, you need the s3:GetObject
// and s3:GetObjectAttributes permissions.
//
// For more information, see [Specifying Permissions in a Policy]in the Amazon S3 User Guide.
//
// If the object that you request does not exist, the error Amazon S3 returns
//
// depends on whether you also have the s3:ListBucket permission.
//
// - If you have the s3:ListBucket permission on the bucket, Amazon S3 returns an
// HTTP status code 404 Not Found ("no such key") error.
//
// - If you don't have the s3:ListBucket permission, Amazon S3 returns an HTTP
// status code 403 Forbidden ("access denied") error.
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
// If the object is encrypted with SSE-KMS, you must also have the
//
// kms:GenerateDataKey and kms:Decrypt permissions in IAM identity-based policies
// and KMS key policies for the KMS key.
//
// Encryption Encryption request headers, like x-amz-server-side-encryption ,
// should not be sent for HEAD requests if your object uses server-side encryption
// with Key Management Service (KMS) keys (SSE-KMS), dual-layer server-side
// encryption with Amazon Web Services KMS keys (DSSE-KMS), or server-side
// encryption with Amazon S3 managed encryption keys (SSE-S3). The
// x-amz-server-side-encryption header is used when you PUT an object to S3 and
// want to specify the encryption method. If you include this header in a GET
// request for an object that uses these types of keys, you’ll get an HTTP 400 Bad
// Request error. It's because the encryption method can't be changed when you
// retrieve the object.
//
// If you encrypted an object when you stored the object in Amazon S3 by using
// server-side encryption with customer-provided encryption keys (SSE-C), then when
// you retrieve the metadata from the object, you must use the following headers.
// These headers provide the server with the encryption key required to retrieve
// the object's metadata. The headers are:
//
// - x-amz-server-side-encryption-customer-algorithm
//
// - x-amz-server-side-encryption-customer-key
//
// - x-amz-server-side-encryption-customer-key-MD5
//
// For more information about SSE-C, see [Server-Side Encryption (Using Customer-Provided Encryption Keys)] in the Amazon S3 User Guide.
//
// Directory bucket permissions - For directory buckets, there are only two
// supported options for server-side encryption: server-side encryption with Amazon
// S3 managed keys (SSE-S3) ( AES256 ) and server-side encryption with KMS keys
// (SSE-KMS) ( aws:kms ). We recommend that the bucket's default encryption uses
// the desired encryption configuration and you don't override the bucket default
// encryption in your CreateSession requests or PUT object requests. Then, new
// objects are automatically encrypted with the desired encryption settings. For
// more information, see [Protecting data with server-side encryption]in the Amazon S3 User Guide. For more information about
// the encryption overriding behaviors in directory buckets, see [Specifying server-side encryption with KMS for new object uploads].
//
// Versioning  Directory buckets - S3 Versioning isn't enabled and supported for
// directory buckets. For this API operation, only the null value of the version
// ID is supported by directory buckets. You can only specify null to the versionId
// query parameter in the request.
//
// Conditional request headers Consider the following when using request headers:
//
// - If both of the If-Match and If-Unmodified-Since headers are present in the
// request as follows, then Amazon S3 returns the HTTP status code 200 OK and the
// data requested:
//
// - If-Match condition evaluates to true .
//
// - If-Unmodified-Since condition evaluates to false .
//
// For more information about conditional requests, see [RFC 7232].
//
// - If both of the If-None-Match and If-Modified-Since headers are present in
// the request as follows, then Amazon S3 returns the HTTP status code 304 Not
// Modified :
//
// - If-None-Match condition evaluates to false .
//
// - If-Modified-Since condition evaluates to true .
//
// For more information about conditional requests, see [RFC 7232].
//
// HTTP Host header syntax  Directory buckets - The HTTP Host header syntax is
// Bucket-name.s3express-zone-id.region-code.amazonaws.com .
//
// The following actions are related to GetObjectAttributes :
//
// [GetObject]
//
// [GetObjectAcl]
//
// [GetObjectLegalHold]
//
// [GetObjectLockConfiguration]
//
// [GetObjectRetention]
//
// [GetObjectTagging]
//
// [HeadObject]
//
// [ListParts]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Specifying server-side encryption with KMS for new object uploads]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-specifying-kms-encryption.html
// [GetObjectLegalHold]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectLegalHold.html
// [Concepts for directory buckets in Local Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-lzs-for-directory-buckets.html
// [ListParts]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListParts.html
// [Server-Side Encryption (Using Customer-Provided Encryption Keys)]: https://docs.aws.amazon.com/AmazonS3/latest/dev/ServerSideEncryptionCustomerKeys.html
// [GetObjectTagging]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectTagging.html
// [Specifying Permissions in a Policy]: https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html
// [RFC 7232]: https://tools.ietf.org/html/rfc7232
// [HeadObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_HeadObject.html
// [GetObjectLockConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectLockConfiguration.html
// [Protecting data with server-side encryption]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-serv-side-encryption.html
// [GetObjectAcl]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectAcl.html
// [GetObjectRetention]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectRetention.html
// [GetObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObject.html
// [Regional and Zonal endpoints for directory buckets in Availability Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/endpoint-directory-buckets-AZ.html
//
// [CreateSession]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateSession.html
