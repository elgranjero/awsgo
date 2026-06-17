package glacier

// UploadArchive is generated as a reference stub.
// Executable command wiring lives under cmd/glacier.go.
//
// This operation adds an archive to a vault. This is a synchronous operation, and
// for a successful upload, your data is durably persisted. Amazon Glacier returns
// the archive ID in the x-amz-archive-id header of the response.
//
// You must use the archive ID to access your data in Amazon Glacier. After you
// upload an archive, you should save the archive ID returned so that you can
// retrieve or delete the archive later. Besides saving the archive ID, you can
// also index it and give it a friendly name to allow for better searching. You can
// also use the optional archive description field to specify how the archive is
// referred to in an external index of archives, such as you might create in Amazon
// DynamoDB. You can also get the vault inventory to obtain a list of archive IDs
// in a vault. For more information, see InitiateJob.
//
// You must provide a SHA256 tree hash of the data you are uploading. For
// information about computing a SHA256 tree hash, see [Computing Checksums].
//
// You can optionally specify an archive description of up to 1,024 printable
// ASCII characters. You can get the archive description when you either retrieve
// the archive or get the vault inventory. For more information, see InitiateJob. Amazon
// Glacier does not interpret the description in any way. An archive description
// does not need to be unique. You cannot use the description to retrieve or sort
// the archive list.
//
// Archives are immutable. After you upload an archive, you cannot edit the
// archive or its description.
//
// An AWS account has full permission to perform all operations (actions).
// However, AWS Identity and Access Management (IAM) users don't have any
// permissions by default. You must grant them explicit permission to perform
// specific actions. For more information, see [Access Control Using AWS Identity and Access Management (IAM)].
//
// For conceptual information and underlying REST API, see [Uploading an Archive in Amazon Glacier] and [Upload Archive] in the Amazon
// Glacier Developer Guide.
//
// [Uploading an Archive in Amazon Glacier]: https://docs.aws.amazon.com/amazonglacier/latest/dev/uploading-an-archive.html
// [Access Control Using AWS Identity and Access Management (IAM)]: https://docs.aws.amazon.com/amazonglacier/latest/dev/using-iam-with-amazon-glacier.html
// [Upload Archive]: https://docs.aws.amazon.com/amazonglacier/latest/dev/api-archive-post.html
// [Computing Checksums]: https://docs.aws.amazon.com/amazonglacier/latest/dev/checksum-calculations.html
