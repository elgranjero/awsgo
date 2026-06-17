package s3

// CreateMultipartUpload is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// End of support notice: As of October 1, 2025, Amazon S3 has discontinued
// support for Email Grantee Access Control Lists (ACLs). If you attempt to use an
// Email Grantee ACL in a request after October 1, 2025, the request will receive
// an HTTP 405 (Method Not Allowed) error.
//
// This change affects the following Amazon Web Services Regions: US East (N.
// Virginia), US West (N. California), US West (Oregon), Asia Pacific (Singapore),
// Asia Pacific (Sydney), Asia Pacific (Tokyo), Europe (Ireland), and South America
// (São Paulo).
//
// This action initiates a multipart upload and returns an upload ID. This upload
// ID is used to associate all of the parts in the specific multipart upload. You
// specify this upload ID in each of your subsequent upload part requests (see [UploadPart]).
// You also include this upload ID in the final request to either complete or abort
// the multipart upload request. For more information about multipart uploads, see [Multipart Upload Overview]
// in the Amazon S3 User Guide.
//
// After you initiate a multipart upload and upload one or more parts, to stop
// being charged for storing the uploaded parts, you must either complete or abort
// the multipart upload. Amazon S3 frees up the space used to store the parts and
// stops charging you for storing them only after you either complete or abort a
// multipart upload.
//
// If you have configured a lifecycle rule to abort incomplete multipart uploads,
// the created multipart upload must be completed within the number of days
// specified in the bucket lifecycle configuration. Otherwise, the incomplete
// multipart upload becomes eligible for an abort action and Amazon S3 aborts the
// multipart upload. For more information, see [Aborting Incomplete Multipart Uploads Using a Bucket Lifecycle Configuration].
//
// - Directory buckets - S3 Lifecycle is not supported by directory buckets.
//
// - Directory buckets - For directory buckets, you must make requests for this
// API operation to the Zonal endpoint. These endpoints support
// virtual-hosted-style requests in the format
// https://amzn-s3-demo-bucket.s3express-zone-id.region-code.amazonaws.com/key-name
// . Path-style requests are not supported. For more information about endpoints
// in Availability Zones, see [Regional and Zonal endpoints for directory buckets in Availability Zones]in the Amazon S3 User Guide. For more information
// about endpoints in Local Zones, see [Concepts for directory buckets in Local Zones]in the Amazon S3 User Guide.
//
// Request signing For request signing, multipart upload is just a series of
// regular requests. You initiate a multipart upload, send one or more requests to
// upload parts, and then complete the multipart upload process. You sign each
// request individually. There is nothing special about signing multipart upload
// requests. For more information about signing, see [Authenticating Requests (Amazon Web Services Signature Version 4)]in the Amazon S3 User Guide.
//
// Permissions
//
// - General purpose bucket permissions - To perform a multipart upload with
// encryption using an Key Management Service (KMS) KMS key, the requester must
// have permission to the kms:Decrypt and kms:GenerateDataKey actions on the key.
// The requester must also have permissions for the kms:GenerateDataKey action
// for the CreateMultipartUpload API. Then, the requester needs permissions for
// the kms:Decrypt action on the UploadPart and UploadPartCopy APIs. These
// permissions are required because Amazon S3 must decrypt and read data from the
// encrypted file parts before it completes the multipart upload. For more
// information, see [Multipart upload API and permissions]and [Protecting data using server-side encryption with Amazon Web Services KMS]in the Amazon S3 User Guide.
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
// Encryption
//
// - General purpose buckets - Server-side encryption is for data encryption at
// rest. Amazon S3 encrypts your data as it writes it to disks in its data centers
// and decrypts it when you access it. Amazon S3 automatically encrypts all new
// objects that are uploaded to an S3 bucket. When doing a multipart upload, if you
// don't specify encryption information in your request, the encryption setting of
// the uploaded parts is set to the default encryption configuration of the
// destination bucket. By default, all buckets have a base level of encryption
// configuration that uses server-side encryption with Amazon S3 managed keys
// (SSE-S3). If the destination bucket has a default encryption configuration that
// uses server-side encryption with an Key Management Service (KMS) key (SSE-KMS),
// or a customer-provided encryption key (SSE-C), Amazon S3 uses the corresponding
// KMS key, or a customer-provided key to encrypt the uploaded parts. When you
// perform a CreateMultipartUpload operation, if you want to use a different type
// of encryption setting for the uploaded parts, you can request that Amazon S3
// encrypts the object with a different encryption key (such as an Amazon S3
// managed key, a KMS key, or a customer-provided key). When the encryption setting
// in your request is different from the default encryption configuration of the
// destination bucket, the encryption setting in your request takes precedence. If
// you choose to provide your own encryption key, the request headers you provide
// in [UploadPart]and [UploadPartCopy]requests must match the headers you used in the CreateMultipartUpload
// request.
//
// - Use KMS keys (SSE-KMS) that include the Amazon Web Services managed key (
// aws/s3 ) and KMS customer managed keys stored in Key Management Service (KMS)
// – If you want Amazon Web Services to manage the keys used to encrypt data,
// specify the following headers in the request.
//
// - x-amz-server-side-encryption
//
// - x-amz-server-side-encryption-aws-kms-key-id
//
// - x-amz-server-side-encryption-context
//
// - If you specify x-amz-server-side-encryption:aws:kms , but don't provide
// x-amz-server-side-encryption-aws-kms-key-id , Amazon S3 uses the Amazon Web
// Services managed key ( aws/s3 key) in KMS to protect the data.
//
// - To perform a multipart upload with encryption by using an Amazon Web
// Services KMS key, the requester must have permission to the kms:Decrypt and
// kms:GenerateDataKey* actions on the key. These permissions are required
// because Amazon S3 must decrypt and read data from the encrypted file parts
// before it completes the multipart upload. For more information, see [Multipart upload API and permissions]and [Protecting data using server-side encryption with Amazon Web Services KMS]in
// the Amazon S3 User Guide.
//
// - If your Identity and Access Management (IAM) user or role is in the same
// Amazon Web Services account as the KMS key, then you must have these permissions
// on the key policy. If your IAM user or role is in a different account from the
// key, then you must have the permissions on both the key policy and your IAM user
// or role.
//
// - All GET and PUT requests for an object protected by KMS fail if you don't
// make them by using Secure Sockets Layer (SSL), Transport Layer Security (TLS),
// or Signature Version 4. For information about configuring any of the officially
// supported Amazon Web Services SDKs and Amazon Web Services CLI, see [Specifying the Signature Version in Request Authentication]in the
// Amazon S3 User Guide.
//
// For more information about server-side encryption with KMS keys (SSE-KMS), see [Protecting Data Using Server-Side Encryption with KMS keys]
//
// in the Amazon S3 User Guide.
//
// - Use customer-provided encryption keys (SSE-C) – If you want to manage your
// own encryption keys, provide all the following headers in the request.
//
// - x-amz-server-side-encryption-customer-algorithm
//
// - x-amz-server-side-encryption-customer-key
//
// - x-amz-server-side-encryption-customer-key-MD5
//
// For more information about server-side encryption with customer-provided
//
// encryption keys (SSE-C), see [Protecting data using server-side encryption with customer-provided encryption keys (SSE-C)]in the Amazon S3 User Guide.
//
// - Directory buckets - For directory buckets, there are only two supported
// options for server-side encryption: server-side encryption with Amazon S3
// managed keys (SSE-S3) ( AES256 ) and server-side encryption with KMS keys
// (SSE-KMS) ( aws:kms ). We recommend that the bucket's default encryption uses
// the desired encryption configuration and you don't override the bucket default
// encryption in your CreateSession requests or PUT object requests. Then, new
// objects are automatically encrypted with the desired encryption settings. For
// more information, see [Protecting data with server-side encryption]in the Amazon S3 User Guide. For more information about
// the encryption overriding behaviors in directory buckets, see [Specifying server-side encryption with KMS for new object uploads].
//
// In the Zonal endpoint API calls (except [CopyObject]and [UploadPartCopy]) using the REST API, the
//
// encryption request headers must match the encryption settings that are specified
// in the CreateSession request. You can't override the values of the encryption
// settings ( x-amz-server-side-encryption ,
// x-amz-server-side-encryption-aws-kms-key-id ,
// x-amz-server-side-encryption-context , and
// x-amz-server-side-encryption-bucket-key-enabled ) that are specified in the
// CreateSession request. You don't need to explicitly specify these encryption
// settings values in Zonal endpoint API calls, and Amazon S3 will use the
// encryption settings values from the CreateSession request to protect new
// objects in the directory bucket.
//
// When you use the CLI or the Amazon Web Services SDKs, for CreateSession , the
//
// session token refreshes automatically to avoid service interruptions when a
// session expires. The CLI or the Amazon Web Services SDKs use the bucket's
// default encryption configuration for the CreateSession request. It's not
// supported to override the encryption settings values in the CreateSession
// request. So in the Zonal endpoint API calls (except [CopyObject]and [UploadPartCopy]), the encryption
// request headers must match the default encryption configuration of the directory
// bucket.
//
// For directory buckets, when you perform a CreateMultipartUpload operation and an
//
// UploadPartCopy operation, the request headers you provide in the
// CreateMultipartUpload request must match the default encryption configuration
// of the destination bucket.
//
// HTTP Host header syntax  Directory buckets - The HTTP Host header syntax is
// Bucket-name.s3express-zone-id.region-code.amazonaws.com .
//
// The following operations are related to CreateMultipartUpload :
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
// [Concepts for directory buckets in Local Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-lzs-for-directory-buckets.html
// [ListParts]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListParts.html
// [UploadPart]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_UploadPart.html
// [Protecting Data Using Server-Side Encryption with KMS keys]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/UsingKMSEncryption.html
// [Specifying the Signature Version in Request Authentication]: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingAWSSDK.html#specify-signature-version
// [Aborting Incomplete Multipart Uploads Using a Bucket Lifecycle Configuration]: https://docs.aws.amazon.com/AmazonS3/latest/dev/mpuoverview.html#mpu-abort-incomplete-mpu-lifecycle-config
// [CopyObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CopyObject.html
// [CreateSession]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateSession.html
// [Multipart upload API and permissions]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/mpuoverview.html#mpuAndPermissions
// [UploadPartCopy]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_UploadPartCopy.html
// [CompleteMultipartUpload]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CompleteMultipartUpload.html
// [Authenticating Requests (Amazon Web Services Signature Version 4)]: https://docs.aws.amazon.com/AmazonS3/latest/API/sig-v4-authenticating-requests.html
// [AbortMultipartUpload]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_AbortMultipartUpload.html
// [Multipart Upload Overview]: https://docs.aws.amazon.com/AmazonS3/latest/dev/mpuoverview.html
// [Protecting data using server-side encryption with Amazon Web Services KMS]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/UsingKMSEncryption.html
// [ListMultipartUploads]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListMultipartUploads.html
// [Regional and Zonal endpoints for directory buckets in Availability Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/endpoint-directory-buckets-AZ.html
//
// [Specifying server-side encryption with KMS for new object uploads]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-specifying-kms-encryption.html
// [Protecting data using server-side encryption with customer-provided encryption keys (SSE-C)]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/ServerSideEncryptionCustomerKeys.html
// [Protecting data with server-side encryption]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-serv-side-encryption.html
