package glacier

// ListMultipartUploads is generated as a reference stub.
// Executable command wiring lives under cmd/glacier.go.
//
// This operation lists in-progress multipart uploads for the specified vault. An
// in-progress multipart upload is a multipart upload that has been initiated by an
// InitiateMultipartUploadrequest, but has not yet been completed or aborted. The list returned in the
// List Multipart Upload response has no guaranteed order.
//
// The List Multipart Uploads operation supports pagination. By default, this
// operation returns up to 50 multipart uploads in the response. You should always
// check the response for a marker at which to continue the list; if there are no
// more items the marker is null . To return a list of multipart uploads that
// begins at a specific upload, set the marker request parameter to the value you
// obtained from a previous List Multipart Upload request. You can also limit the
// number of uploads returned in the response by specifying the limit parameter in
// the request.
//
// Note the difference between this operation and listing parts (ListParts ). The List
// Multipart Uploads operation lists all multipart uploads for a vault and does not
// require a multipart upload ID. The List Parts operation requires a multipart
// upload ID since parts are associated with a single upload.
//
// An AWS account has full permission to perform all operations (actions).
// However, AWS Identity and Access Management (IAM) users don't have any
// permissions by default. You must grant them explicit permission to perform
// specific actions. For more information, see [Access Control Using AWS Identity and Access Management (IAM)].
//
// For conceptual information and the underlying REST API, see [Working with Archives in Amazon Glacier] and [List Multipart Uploads] in the Amazon
// Glacier Developer Guide.
//
// [Access Control Using AWS Identity and Access Management (IAM)]: https://docs.aws.amazon.com/amazonglacier/latest/dev/using-iam-with-amazon-glacier.html
// [List Multipart Uploads]: https://docs.aws.amazon.com/amazonglacier/latest/dev/api-multipart-list-uploads.html
// [Working with Archives in Amazon Glacier]: https://docs.aws.amazon.com/amazonglacier/latest/dev/working-with-archives.html
