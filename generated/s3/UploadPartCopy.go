package s3

// UploadPartCopy is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// Uploads a part by copying data from an existing object as data source. To
// specify the data source, you add the request header x-amz-copy-source in your
// request. To specify a byte range, you add the request header
// x-amz-copy-source-range in your request.
//
// For information about maximum and minimum part sizes and other multipart upload
// specifications, see [Multipart upload limits]in the Amazon S3 User Guide.
//
// Instead of copying data from an existing object as part data, you might use the [UploadPart]
// action to upload new data as a part of an object in your request.
//
// You must initiate a multipart upload before you can upload any part. In
// response to your initiate request, Amazon S3 returns the upload ID, a unique
// identifier that you must include in your upload part request.
//
// For conceptual information about multipart uploads, see [Uploading Objects Using Multipart Upload] in the Amazon S3 User
// Guide. For information about copying objects using a single atomic action vs. a
// multipart upload, see [Operations on Objects]in the Amazon S3 User Guide.
//
// Directory buckets - For directory buckets, you must make requests for this API
// operation to the Zonal endpoint. These endpoints support virtual-hosted-style
// requests in the format
// https://amzn-s3-demo-bucket.s3express-zone-id.region-code.amazonaws.com/key-name
// . Path-style requests are not supported. For more information about endpoints
// in Availability Zones, see [Regional and Zonal endpoints for directory buckets in Availability Zones]in the Amazon S3 User Guide. For more information
// about endpoints in Local Zones, see [Concepts for directory buckets in Local Zones]in the Amazon S3 User Guide.
//
// Authentication and authorization All UploadPartCopy requests must be
// authenticated and signed by using IAM credentials (access key ID and secret
// access key for the IAM identities). All headers with the x-amz- prefix,
// including x-amz-copy-source , must be signed. For more information, see [REST Authentication].
//
// Directory buckets - You must use IAM credentials to authenticate and authorize
// your access to the UploadPartCopy API operation, instead of using the temporary
// security credentials through the CreateSession API operation.
//
// Amazon Web Services CLI or SDKs handles authentication and authorization on
// your behalf.
//
// Permissions You must have READ access to the source object and WRITE access to
// the destination bucket.
//
// - General purpose bucket permissions - You must have the permissions in a
// policy based on the bucket types of your source bucket and destination bucket in
// an UploadPartCopy operation.
//
// - If the source object is in a general purpose bucket, you must have the
// s3:GetObject permission to read the source object that is being copied.
//
// - If the destination bucket is a general purpose bucket, you must have the
// s3:PutObject permission to write the object copy to the destination bucket.
//
// - To perform a multipart upload with encryption using an Key Management
// Service key, the requester must have permission to the kms:Decrypt and
// kms:GenerateDataKey actions on the key. The requester must also have
// permissions for the kms:GenerateDataKey action for the CreateMultipartUpload
// API. Then, the requester needs permissions for the kms:Decrypt action on the
// UploadPart and UploadPartCopy APIs. These permissions are required because
// Amazon S3 must decrypt and read data from the encrypted file parts before it
// completes the multipart upload. For more information about KMS permissions, see [Protecting data using server-side encryption with KMS]
// in the Amazon S3 User Guide. For information about the permissions required to
// use the multipart upload API, see [Multipart upload and permissions]and [Multipart upload API and permissions]in the Amazon S3 User Guide.
//
// - Directory bucket permissions - You must have permissions in a bucket policy
// or an IAM identity-based policy based on the source and destination bucket types
// in an UploadPartCopy operation.
//
// - If the source object that you want to copy is in a directory bucket, you
// must have the s3express:CreateSession permission in the Action element of a
// policy to read the object. By default, the session is in the ReadWrite mode.
// If you want to restrict the access, you can explicitly set the
// s3express:SessionMode condition key to ReadOnly on the copy source bucket.
//
// - If the copy destination is a directory bucket, you must have the
// s3express:CreateSession permission in the Action element of a policy to write
// the object to the destination. The s3express:SessionMode condition key cannot
// be set to ReadOnly on the copy destination.
//
// If the object is encrypted with SSE-KMS, you must also have the
//
// kms:GenerateDataKey and kms:Decrypt permissions in IAM identity-based policies
// and KMS key policies for the KMS key.
//
// For example policies, see [Example bucket policies for S3 Express One Zone]and [Amazon Web Services Identity and Access Management (IAM) identity-based policies for S3 Express One Zone]in the Amazon S3 User Guide.
//
// Encryption
// - General purpose buckets - For information about using server-side
// encryption with customer-provided encryption keys with the UploadPartCopy
// operation, see [CopyObject]and [UploadPart].
//
// If you have server-side encryption with customer-provided keys (SSE-C) blocked
//
// for your general purpose bucket, you will get an HTTP 403 Access Denied error
// when you specify the SSE-C request headers while writing new data to your
// bucket. For more information, see [Blocking or unblocking SSE-C for a general purpose bucket].
//
// - Directory buckets - For directory buckets, there are only two supported
// options for server-side encryption: server-side encryption with Amazon S3
// managed keys (SSE-S3) ( AES256 ) and server-side encryption with KMS keys
// (SSE-KMS) ( aws:kms ). For more information, see [Protecting data with server-side encryption]in the Amazon S3 User Guide.
//
// For directory buckets, when you perform a CreateMultipartUpload operation and an
//
// UploadPartCopy operation, the request headers you provide in the
// CreateMultipartUpload request must match the default encryption configuration
// of the destination bucket.
//
// S3 Bucket Keys aren't supported, when you copy SSE-KMS encrypted objects from
//
// general purpose buckets to directory buckets, from directory buckets to general
// purpose buckets, or between directory buckets, through [UploadPartCopy]. In this case, Amazon
// S3 makes a call to KMS every time a copy request is made for a KMS-encrypted
// object.
//
// Special errors
//
// - Error Code: NoSuchUpload
//
// - Description: The specified multipart upload does not exist. The upload ID
// might be invalid, or the multipart upload might have been aborted or completed.
//
// - HTTP Status Code: 404 Not Found
//
// - Error Code: InvalidRequest
//
// - Description: The specified copy source is not supported as a byte-range
// copy source.
//
// - HTTP Status Code: 400 Bad Request
//
// HTTP Host header syntax  Directory buckets - The HTTP Host header syntax is
// Bucket-name.s3express-zone-id.region-code.amazonaws.com .
//
// The following operations are related to UploadPartCopy :
//
// [CreateMultipartUpload]
//
// [UploadPart]
//
// [CompleteMultipartUpload]
//
// [AbortMultipartUpload]
//
// [ListParts]
//
// [ListMultipartUploads]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Uploading Objects Using Multipart Upload]: https://docs.aws.amazon.com/AmazonS3/latest/dev/uploadobjusingmpu.html
// [Concepts for directory buckets in Local Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-lzs-for-directory-buckets.html
// [ListParts]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListParts.html
// [UploadPart]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_UploadPart.html
// [Protecting data using server-side encryption with KMS]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/UsingKMSEncryption.html
// [CopyObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CopyObject.html
// [Multipart upload and permissions]: https://docs.aws.amazon.com/AmazonS3/latest/dev/mpuAndPermissions.html
// [Multipart upload API and permissions]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/mpuoverview.html#mpuAndPermissions
// [CompleteMultipartUpload]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CompleteMultipartUpload.html
// [CreateMultipartUpload]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateMultipartUpload.html
// [Multipart upload limits]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/qfacts.html
// [Amazon Web Services Identity and Access Management (IAM) identity-based policies for S3 Express One Zone]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-security-iam-identity-policies.html
// [AbortMultipartUpload]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_AbortMultipartUpload.html
// [REST Authentication]: https://docs.aws.amazon.com/AmazonS3/latest/dev/RESTAuthentication.html
// [Example bucket policies for S3 Express One Zone]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-security-iam-example-bucket-policies.html
// [Operations on Objects]: https://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectOperations.html
// [ListMultipartUploads]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListMultipartUploads.html
// [Regional and Zonal endpoints for directory buckets in Availability Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/endpoint-directory-buckets-AZ.html
//
// [Blocking or unblocking SSE-C for a general purpose bucket]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/blocking-unblocking-s3-c-encryption-gpb.html
// [UploadPartCopy]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_UploadPartCopy.html
// [Protecting data with server-side encryption]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-serv-side-encryption.html
