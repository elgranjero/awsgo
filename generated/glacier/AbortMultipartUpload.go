package glacier

// AbortMultipartUpload is generated as a reference stub.
// Executable command wiring lives under cmd/glacier.go.
//
// This operation aborts a multipart upload identified by the upload ID.
//
// After the Abort Multipart Upload request succeeds, you cannot upload any more
// parts to the multipart upload or complete the multipart upload. Aborting a
// completed upload fails. However, aborting an already-aborted upload will
// succeed, for a short time. For more information about uploading a part and
// completing a multipart upload, see UploadMultipartPartand CompleteMultipartUpload.
//
// This operation is idempotent.
//
// An AWS account has full permission to perform all operations (actions).
// However, AWS Identity and Access Management (IAM) users don't have any
// permissions by default. You must grant them explicit permission to perform
// specific actions. For more information, see [Access Control Using AWS Identity and Access Management (IAM)].
//
// For conceptual information and underlying REST API, see [Working with Archives in Amazon Glacier] and [Abort Multipart Upload] in the Amazon
// Glacier Developer Guide.
//
// [Abort Multipart Upload]: https://docs.aws.amazon.com/amazonglacier/latest/dev/api-multipart-abort-upload.html
// [Access Control Using AWS Identity and Access Management (IAM)]: https://docs.aws.amazon.com/amazonglacier/latest/dev/using-iam-with-amazon-glacier.html
// [Working with Archives in Amazon Glacier]: https://docs.aws.amazon.com/amazonglacier/latest/dev/working-with-archives.html
