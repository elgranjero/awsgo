package s3

// GetObject is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// Retrieves an object from Amazon S3.
//
// In the GetObject request, specify the full key name for the object.
//
// General purpose buckets - Both the virtual-hosted-style requests and the
// path-style requests are supported. For a virtual hosted-style request example,
// if you have the object photos/2006/February/sample.jpg , specify the object key
// name as /photos/2006/February/sample.jpg . For a path-style request example, if
// you have the object photos/2006/February/sample.jpg in the bucket named
// examplebucket , specify the object key name as
// /examplebucket/photos/2006/February/sample.jpg . For more information about
// request types, see [HTTP Host Header Bucket Specification]in the Amazon S3 User Guide.
//
// Directory buckets - Only virtual-hosted-style requests are supported. For a
// virtual hosted-style request example, if you have the object
// photos/2006/February/sample.jpg in the bucket named
// amzn-s3-demo-bucket--usw2-az1--x-s3 , specify the object key name as
// /photos/2006/February/sample.jpg . Also, when you make requests to this API
// operation, your requests are sent to the Zonal endpoint. These endpoints support
// virtual-hosted-style requests in the format
// https://bucket-name.s3express-zone-id.region-code.amazonaws.com/key-name .
// Path-style requests are not supported. For more information about endpoints in
// Availability Zones, see [Regional and Zonal endpoints for directory buckets in Availability Zones]in the Amazon S3 User Guide. For more information about
// endpoints in Local Zones, see [Concepts for directory buckets in Local Zones]in the Amazon S3 User Guide.
//
// Permissions
// - General purpose bucket permissions - You must have the required permissions
// in a policy. To use GetObject , you must have the READ access to the object
// (or version). If you grant READ access to the anonymous user, the GetObject
// operation returns the object without using an authorization header. For more
// information, see [Specifying permissions in a policy]in the Amazon S3 User Guide.
//
// If you include a versionId in your request header, you must have the
//
// s3:GetObjectVersion permission to access a specific version of an object. The
// s3:GetObject permission is not required in this scenario.
//
// If you request the current version of an object without a specific versionId in
//
// the request header, only the s3:GetObject permission is required. The
// s3:GetObjectVersion permission is not required in this scenario.
//
// If the object that you request doesn’t exist, the error that Amazon S3 returns
//
// depends on whether you also have the s3:ListBucket permission.
//
// - If you have the s3:ListBucket permission on the bucket, Amazon S3 returns an
// HTTP status code 404 Not Found error.
//
// - If you don’t have the s3:ListBucket permission, Amazon S3 returns an HTTP
// status code 403 Access Denied error.
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
// If the object is encrypted using SSE-KMS, you must also have the
//
// kms:GenerateDataKey and kms:Decrypt permissions in IAM identity-based policies
// and KMS key policies for the KMS key.
//
// Storage classes If the object you are retrieving is stored in the S3 Glacier
// Flexible Retrieval storage class, the S3 Glacier Deep Archive storage class, the
// S3 Intelligent-Tiering Archive Access tier, or the S3 Intelligent-Tiering Deep
// Archive Access tier, before you can retrieve the object you must first restore a
// copy using [RestoreObject]. Otherwise, this operation returns an InvalidObjectState error. For
// information about restoring archived objects, see [Restoring Archived Objects]in the Amazon S3 User Guide.
//
// Directory buckets - Directory buckets only support EXPRESS_ONEZONE (the S3
// Express One Zone storage class) in Availability Zones and ONEZONE_IA (the S3
// One Zone-Infrequent Access storage class) in Dedicated Local Zones. Unsupported
// storage class values won't write a destination object and will respond with the
// HTTP status code 400 Bad Request .
//
// Encryption Encryption request headers, like x-amz-server-side-encryption ,
// should not be sent for the GetObject requests, if your object uses server-side
// encryption with Amazon S3 managed encryption keys (SSE-S3), server-side
// encryption with Key Management Service (KMS) keys (SSE-KMS), or dual-layer
// server-side encryption with Amazon Web Services KMS keys (DSSE-KMS). If you
// include the header in your GetObject requests for the object that uses these
// types of keys, you’ll get an HTTP 400 Bad Request error.
//
// Directory buckets - For directory buckets, there are only two supported options
// for server-side encryption: SSE-S3 and SSE-KMS. SSE-C isn't supported. For more
// information, see [Protecting data with server-side encryption]in the Amazon S3 User Guide.
//
// Overriding response header values through the request There are times when you
// want to override certain response header values of a GetObject response. For
// example, you might override the Content-Disposition response header value
// through your GetObject request.
//
// You can override values for a set of response headers. These modified response
// header values are included only in a successful response, that is, when the HTTP
// status code 200 OK is returned. The headers you can override using the
// following query parameters in the request are a subset of the headers that
// Amazon S3 accepts when you create an object.
//
// The response headers that you can override for the GetObject response are
// Cache-Control , Content-Disposition , Content-Encoding , Content-Language ,
// Content-Type , and Expires .
//
// To override values for a set of response headers in the GetObject response, you
// can use the following query parameters in the request.
//
// - response-cache-control
//
// - response-content-disposition
//
// - response-content-encoding
//
// - response-content-language
//
// - response-content-type
//
// - response-expires
//
// When you use these parameters, you must sign the request by using either an
// Authorization header or a presigned URL. These parameters cannot be used with an
// unsigned (anonymous) request.
//
// HTTP Host header syntax  Directory buckets - The HTTP Host header syntax is
// Bucket-name.s3express-zone-id.region-code.amazonaws.com .
//
// The following operations are related to GetObject :
//
// [ListBuckets]
//
// [GetObjectAcl]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Concepts for directory buckets in Local Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-lzs-for-directory-buckets.html
// [RestoreObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_RestoreObject.html
// [Protecting data with server-side encryption]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-serv-side-encryption.html
// [ListBuckets]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListBuckets.html
// [HTTP Host Header Bucket Specification]: https://docs.aws.amazon.com/AmazonS3/latest/dev/VirtualHosting.html#VirtualHostingSpecifyBucket
// [Restoring Archived Objects]: https://docs.aws.amazon.com/AmazonS3/latest/dev/restoring-objects.html
// [GetObjectAcl]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectAcl.html
// [Specifying permissions in a policy]: https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html
// [Regional and Zonal endpoints for directory buckets in Availability Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/endpoint-directory-buckets-AZ.html
//
// [CreateSession]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateSession.html
