package glacier

// DeleteArchive is generated as a reference stub.
// Executable command wiring lives under cmd/glacier.go.
//
// This operation deletes an archive from a vault. Subsequent requests to initiate
// a retrieval of this archive will fail. Archive retrievals that are in progress
// for this archive ID may or may not succeed according to the following scenarios:
//
// - If the archive retrieval job is actively preparing the data for download
// when Amazon Glacier receives the delete archive request, the archival retrieval
// operation might fail.
//
// - If the archive retrieval job has successfully prepared the archive for
// download when Amazon Glacier receives the delete archive request, you will be
// able to download the output.
//
// This operation is idempotent. Attempting to delete an already-deleted archive
// does not result in an error.
//
// An AWS account has full permission to perform all operations (actions).
// However, AWS Identity and Access Management (IAM) users don't have any
// permissions by default. You must grant them explicit permission to perform
// specific actions. For more information, see [Access Control Using AWS Identity and Access Management (IAM)].
//
// For conceptual information and underlying REST API, see [Deleting an Archive in Amazon Glacier] and [Delete Archive] in the Amazon
// Glacier Developer Guide.
//
// [Delete Archive]: https://docs.aws.amazon.com/amazonglacier/latest/dev/api-archive-delete.html
// [Access Control Using AWS Identity and Access Management (IAM)]: https://docs.aws.amazon.com/amazonglacier/latest/dev/using-iam-with-amazon-glacier.html
// [Deleting an Archive in Amazon Glacier]: https://docs.aws.amazon.com/amazonglacier/latest/dev/deleting-an-archive.html
