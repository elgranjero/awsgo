package s3

// UploadPart is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// Uploads a part in a multipart upload.
//
// In this operation, you provide new data as a part of an object in your request.
// However, you have an option to specify your existing Amazon S3 object as a data
// source for the part you are uploading. To upload a part from an existing object,
// you use the [UploadPartCopy]operation.
//
// You must initiate a multipart upload (see [CreateMultipartUpload]) before you can upload any part. In
// response to your initiate request, Amazon S3 returns an upload ID, a unique
// identifier that you must include in your upload part request.
//
// Part numbers can be any number from 1 to 10,000, inclusive. A part number
// uniquely identifies a part and also defines its position within the object being
// created. If you upload a new part using the same part number that was used with
// a previous part, the previously uploaded part is overwritten.
//
// For information about maximum and minimum part sizes and other multipart upload
// specifications, see [Multipart upload limits]in the Amazon S3 User Guide.
//
// After you initiate multipart upload and upload one or more parts, you must
// either complete or abort multipart upload in order to stop getting charged for
// storage of the uploaded parts. Only after you either complete or abort multipart
// upload, Amazon S3 frees up the parts storage and stops charging you for the
// parts storage.
//
// For more information on multipart uploads, go to [Multipart Upload Overview] in the Amazon S3 User Guide .
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
// - General purpose bucket permissions - To perform a multipart upload with
// encryption using an Key Management Service key, the requester must have
// permission to the kms:Decrypt and kms:GenerateDataKey actions on the key. The
// requester must also have permissions for the kms:GenerateDataKey action for
// the CreateMultipartUpload API. Then, the requester needs permissions for the
// kms:Decrypt action on the UploadPart and UploadPartCopy APIs.
//
// These permissions are required because Amazon S3 must decrypt and read data
//
// from the encrypted file parts before it completes the multipart upload. For more
// information about KMS permissions, see [Protecting data using server-side encryption with KMS]in the Amazon S3 User Guide. For
// information about the permissions required to use the multipart upload API, see [Multipart upload and permissions]
// and [Multipart upload API and permissions]in the Amazon S3 User Guide.
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
// Data integrity  General purpose bucket - To ensure that data is not corrupted
// traversing the network, specify the Content-MD5 header in the upload part
// request. Amazon S3 checks the part data against the provided MD5 value. If they
// do not match, Amazon S3 returns an error. If the upload request is signed with
// Signature Version 4, then Amazon Web Services S3 uses the x-amz-content-sha256
// header as a checksum instead of Content-MD5 . For more information see [Authenticating Requests: Using the Authorization Header (Amazon Web Services Signature Version 4)].
//
// Directory buckets - MD5 is not supported by directory buckets. You can use
// checksum algorithms to check object integrity.
//
// Encryption
// - General purpose bucket - Server-side encryption is for data encryption at
// rest. Amazon S3 encrypts your data as it writes it to disks in its data centers
// and decrypts it when you access it. You have mutually exclusive options to
// protect data using server-side encryption in Amazon S3, depending on how you
// choose to manage the encryption keys. Specifically, the encryption key options
// are Amazon S3 managed keys (SSE-S3), Amazon Web Services KMS keys (SSE-KMS), and
// Customer-Provided Keys (SSE-C). Amazon S3 encrypts data with server-side
// encryption using Amazon S3 managed keys (SSE-S3) by default. You can optionally
// tell Amazon S3 to encrypt data at rest using server-side encryption with other
// key options. The option you use depends on whether you want to use KMS keys
// (SSE-KMS) or provide your own encryption key (SSE-C).
//
// Server-side encryption is supported by the S3 Multipart Upload operations.
//
// Unless you are using a customer-provided encryption key (SSE-C), you don't need
// to specify the encryption parameters in each UploadPart request. Instead, you
// only need to specify the server-side encryption parameters in the initial
// Initiate Multipart request. For more information, see [CreateMultipartUpload].
//
// If you have server-side encryption with customer-provided keys (SSE-C) blocked
//
// for your general purpose bucket, you will get an HTTP 403 Access Denied error
// when you specify the SSE-C request headers while writing new data to your
// bucket. For more information, see [Blocking or unblocking SSE-C for a general purpose bucket].
//
// If you request server-side encryption using a customer-provided encryption key
//
// (SSE-C) in your initiate multipart upload request, you must provide identical
// encryption information in each part upload using the following request headers.
//
// - x-amz-server-side-encryption-customer-algorithm
//
// - x-amz-server-side-encryption-customer-key
//
// - x-amz-server-side-encryption-customer-key-MD5
//
// For more information, see [Using Server-Side Encryption]in the Amazon S3 User Guide.
//
// - Directory buckets - For directory buckets, there are only two supported
// options for server-side encryption: server-side encryption with Amazon S3
// managed keys (SSE-S3) ( AES256 ) and server-side encryption with KMS keys
// (SSE-KMS) ( aws:kms ).
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
// - SOAP Fault Code Prefix: Client
//
// HTTP Host header syntax  Directory buckets - The HTTP Host header syntax is
// Bucket-name.s3express-zone-id.region-code.amazonaws.com .
//
// The following operations are related to UploadPart :
//
// [CreateMultipartUpload]
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
// [Concepts for directory buckets in Local Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-lzs-for-directory-buckets.html
// [ListParts]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListParts.html
// [Authenticating Requests: Using the Authorization Header (Amazon Web Services Signature Version 4)]: https://docs.aws.amazon.com/AmazonS3/latest/API/sigv4-auth-using-authorization-header.html
// [UploadPartCopy]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_UploadPartCopy.html
// [CompleteMultipartUpload]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CompleteMultipartUpload.html
// [CreateMultipartUpload]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateMultipartUpload.html
// [Using Server-Side Encryption]: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingServerSideEncryption.html
// [Multipart upload limits]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/qfacts.html
// [AbortMultipartUpload]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_AbortMultipartUpload.html
// [Multipart Upload Overview]: https://docs.aws.amazon.com/AmazonS3/latest/dev/mpuoverview.html
// [ListMultipartUploads]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListMultipartUploads.html
// [Regional and Zonal endpoints for directory buckets in Availability Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/endpoint-directory-buckets-AZ.html
//
// [Protecting data using server-side encryption with KMS]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/UsingKMSEncryption.html
// [Multipart upload and permissions]: https://docs.aws.amazon.com/AmazonS3/latest/dev/mpuAndPermissions.html
// [CreateSession]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateSession.html
// [Multipart upload API and permissions]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/mpuoverview.html#mpuAndPermissions
// [Blocking or unblocking SSE-C for a general purpose bucket]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/blocking-unblocking-s3-c-encryption-gpb.html
