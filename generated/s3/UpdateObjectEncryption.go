package s3

// UpdateObjectEncryption is generated as a reference stub.
// Executable command wiring lives under cmd/s3.go.
//
// This operation is not supported for directory buckets or Amazon S3 on Outposts
// buckets.
//
// Updates the server-side encryption type of an existing encrypted object in a
// general purpose bucket. You can use the UpdateObjectEncryption operation to
// change encrypted objects from server-side encryption with Amazon S3 managed keys
// (SSE-S3) to server-side encryption with Key Management Service (KMS) keys
// (SSE-KMS), or to apply S3 Bucket Keys. You can also use the
// UpdateObjectEncryption operation to change the customer-managed KMS key used to
// encrypt your data so that you can comply with custom key-rotation standards.
//
// Using the UpdateObjectEncryption operation, you can atomically update the
// server-side encryption type of an existing object in a general purpose bucket
// without any data movement. The UpdateObjectEncryption operation uses envelope
// encryption to re-encrypt the data key used to encrypt and decrypt your object
// with your newly specified server-side encryption type. In other words, when you
// use the UpdateObjectEncryption operation, your data isn't copied, archived
// objects in the S3 Glacier Flexible Retrieval and S3 Glacier Deep Archive storage
// classes aren't restored, and objects in the S3 Intelligent-Tiering storage class
// aren't moved between tiers. Additionally, the UpdateObjectEncryption operation
// preserves all object metadata properties, including the storage class, creation
// date, last modified date, ETag, and checksum properties. For more information,
// see [Updating server-side encryption for existing objects]in the Amazon S3 User Guide.
//
// By default, all UpdateObjectEncryption requests that specify a customer-managed
// KMS key are restricted to KMS keys that are owned by the bucket owner's Amazon
// Web Services account. If you're using Organizations, you can request the ability
// to use KMS keys owned by other member accounts within your organization by
// contacting Amazon Web Services Support.
//
// Source objects that are unencrypted, or encrypted with either dual-layer
// server-side encryption with KMS keys (DSSE-KMS) or server-side encryption with
// customer-provided keys (SSE-C) aren't supported by this operation. Additionally,
// you cannot specify SSE-S3 encryption as the requested new encryption type
// UpdateObjectEncryption request.
//
// Permissions
//
// - To use the UpdateObjectEncryption operation, you must have the following
// permissions:
//
// - s3:PutObject
//
// - s3:UpdateObjectEncryption
//
// - kms:Encrypt
//
// - kms:Decrypt
//
// - kms:GenerateDataKey
//
// - kms:ReEncrypt*
//
// - If you're using Organizations, to use this operation with customer-managed
// KMS keys from other Amazon Web Services accounts within your organization, you
// must have the organizations:DescribeAccount permission.
//
// Errors
//
// - You might receive an InvalidRequest error for several reasons. Depending on
// the reason for the error, you might receive one of the following messages:
//
// - The UpdateObjectEncryption operation doesn't supported unencrypted source
// objects. Only source objects encrypted with SSE-S3 or SSE-KMS are supported.
//
// - The UpdateObjectEncryption operation doesn't support source objects with the
// encryption type DSSE-KMS or SSE-C. Only source objects encrypted with SSE-S3 or
// SSE-KMS are supported.
//
// - The UpdateObjectEncryption operation doesn't support updating the encryption
// type to DSSE-KMS or SSE-C. Modify the request to specify SSE-KMS for the updated
// encryption type, and then try again.
//
// - Requests that modify an object encryption configuration require Amazon Web
// Services Signature Version 4. Modify the request to use Amazon Web Services
// Signature Version 4, and then try again.
//
// - Requests that modify an object encryption configuration require a valid new
// encryption type. Valid values are SSEKMS . Modify the request to specify
// SSE-KMS for the updated encryption type, and then try again.
//
// - Requests that modify an object's encryption type to SSE-KMS require an
// Amazon Web Services KMS key Amazon Resource Name (ARN). Modify the request to
// specify a KMS key ARN, and then try again.
//
// - Requests that modify an object's encryption type to SSE-KMS require a valid
// Amazon Web Services KMS key Amazon Resource Name (ARN). Confirm that you have a
// correctly formatted KMS key ARN in your request, and then try again.
//
// - The BucketKeyEnabled value isn't valid. Valid values are true or false .
// Modify the request to specify a valid value, and then try again.
//
// - You might receive an AccessDenied error for several reasons. Depending on
// the reason for the error, you might receive one of the following messages:
//
// - The Amazon Web Services KMS key in the request must be owned by the same
// account as the bucket. Modify the request to specify a KMS key from the same
// account, and then try again.
//
// - The bucket owner's account was approved to make UpdateObjectEncryption
// requests that use any Amazon Web Services KMS key in their organization, but the
// bucket owner's account isn't part of an organization in Organizations. Make sure
// that the bucket owner's account and the specified KMS key belong to the same
// organization, and then try again.
//
// - The specified Amazon Web Services KMS key must be from the same
// organization in Organizations as the bucket. Specify a KMS key that belongs to
// the same organization as the bucket, and then try again.
//
// - The encryption type for the specified object can’t be updated because that
// object is protected by S3 Object Lock. If the object has a governance-mode
// retention period or a legal hold, you must first remove the Object Lock status
// on the object before you issue your UpdateObjectEncryption request. You can't
// use the UpdateObjectEncryption operation with objects that have an Object Lock
// compliance mode retention period applied to them.
//
// [Updating server-side encryption for existing objects]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/update-sse-encryption.html
