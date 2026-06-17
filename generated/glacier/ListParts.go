package glacier

// ListParts is generated as a reference stub.
// Executable command wiring lives under cmd/glacier.go.
//
// This operation lists the parts of an archive that have been uploaded in a
// specific multipart upload. You can make this request at any time during an
// in-progress multipart upload before you complete the upload (see CompleteMultipartUpload. List Parts
// returns an error for completed uploads. The list returned in the List Parts
// response is sorted by part range.
//
// The List Parts operation supports pagination. By default, this operation
// returns up to 50 uploaded parts in the response. You should always check the
// response for a marker at which to continue the list; if there are no more items
// the marker is null . To return a list of parts that begins at a specific part,
// set the marker request parameter to the value you obtained from a previous List
// Parts request. You can also limit the number of parts returned in the response
// by specifying the limit parameter in the request.
//
// An AWS account has full permission to perform all operations (actions).
// However, AWS Identity and Access Management (IAM) users don't have any
// permissions by default. You must grant them explicit permission to perform
// specific actions. For more information, see [Access Control Using AWS Identity and Access Management (IAM)].
//
// For conceptual information and the underlying REST API, see [Working with Archives in Amazon Glacier] and [List Parts] in the Amazon
// Glacier Developer Guide.
//
// [List Parts]: https://docs.aws.amazon.com/amazonglacier/latest/dev/api-multipart-list-parts.html
// [Access Control Using AWS Identity and Access Management (IAM)]: https://docs.aws.amazon.com/amazonglacier/latest/dev/using-iam-with-amazon-glacier.html
// [Working with Archives in Amazon Glacier]: https://docs.aws.amazon.com/amazonglacier/latest/dev/working-with-archives.html
