package s3

// DeleteObjects is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// This operation enables you to delete multiple objects from a bucket using a
// single HTTP request. If you know the object keys that you want to delete, then
// this operation provides a suitable alternative to sending individual delete
// requests, reducing per-request overhead.
//
// The request can contain a list of up to 1,000 keys that you want to delete. In
// the XML, you provide the object key names, and optionally, version IDs if you
// want to delete a specific version of the object from a versioning-enabled
// bucket. For each key, Amazon S3 performs a delete operation and returns the
// result of that delete, success or failure, in the response. If the object
// specified in the request isn't found, Amazon S3 confirms the deletion by
// returning the result as deleted.
//
// - Directory buckets - S3 Versioning isn't enabled and supported for directory
// buckets.
//
// - Directory buckets - For directory buckets, you must make requests for this
// API operation to the Zonal endpoint. These endpoints support
// virtual-hosted-style requests in the format
// https://amzn-s3-demo-bucket.s3express-zone-id.region-code.amazonaws.com/key-name
// . Path-style requests are not supported. For more information about endpoints
// in Availability Zones, see [Regional and Zonal endpoints for directory buckets in Availability Zones]in the Amazon S3 User Guide. For more information
// about endpoints in Local Zones, see [Concepts for directory buckets in Local Zones]in the Amazon S3 User Guide.
//
// The operation supports two modes for the response: verbose and quiet. By
// default, the operation uses verbose mode in which the response includes the
// result of deletion of each key in your request. In quiet mode the response
// includes only keys where the delete operation encountered an error. For a
// successful deletion in a quiet mode, the operation does not return any
// information about the delete in the response body.
//
// When performing this action on an MFA Delete enabled bucket, that attempts to
// delete any versioned objects, you must include an MFA token. If you do not
// provide one, the entire request will fail, even if there are non-versioned
// objects you are trying to delete. If you provide an invalid token, whether there
// are versioned keys in the request or not, the entire Multi-Object Delete request
// will fail. For information about MFA Delete, see [MFA Delete]in the Amazon S3 User Guide.
//
// Directory buckets - MFA delete is not supported by directory buckets.
//
// Permissions
//
// - General purpose bucket permissions - The following permissions are required
// in your policies when your DeleteObjects request includes specific headers.
//
// - s3:DeleteObject - To delete an object from a bucket, you must always specify
// the s3:DeleteObject permission.
//
// - s3:DeleteObjectVersion - To delete a specific version of an object from a
// versioning-enabled bucket, you must specify the s3:DeleteObjectVersion
// permission.
//
// If the s3:DeleteObject or s3:DeleteObjectVersion permissions are explicitly
//
// denied in your bucket policy, attempts to delete any unversioned objects result
// in a 403 Access Denied error.
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
// Content-MD5 request header
//
// - General purpose bucket - The Content-MD5 request header is required for all
// Multi-Object Delete requests. Amazon S3 uses the header value to ensure that
// your request body has not been altered in transit.
//
// - Directory bucket - The Content-MD5 request header or a additional checksum
// request header (including x-amz-checksum-crc32 , x-amz-checksum-crc32c ,
// x-amz-checksum-sha1 , or x-amz-checksum-sha256 ) is required for all
// Multi-Object Delete requests.
//
// HTTP Host header syntax  Directory buckets - The HTTP Host header syntax is
// Bucket-name.s3express-zone-id.region-code.amazonaws.com .
//
// The following operations are related to DeleteObjects :
//
// [CreateMultipartUpload]
//
// [UploadPart]
//
// [CompleteMultipartUpload]
//
// [ListParts]
//
// [AbortMultipartUpload]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Concepts for directory buckets in Local Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-lzs-for-directory-buckets.html
// [ListParts]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListParts.html
// [AbortMultipartUpload]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_AbortMultipartUpload.html
// [UploadPart]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_UploadPart.html
// [CompleteMultipartUpload]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CompleteMultipartUpload.html
// [Regional and Zonal endpoints for directory buckets in Availability Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/endpoint-directory-buckets-AZ.html
// [MFA Delete]: https://docs.aws.amazon.com/AmazonS3/latest/dev/Versioning.html#MultiFactorAuthenticationDelete
// [CreateMultipartUpload]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateMultipartUpload.html
//
// [CreateSession]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateSession.html
