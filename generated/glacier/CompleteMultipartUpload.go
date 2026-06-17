package glacier

// CompleteMultipartUpload is generated as a reference stub.
// Executable command wiring lives under cmd/glacier.go.
//
// You call this operation to inform Amazon Glacier (Glacier) that all the archive
// parts have been uploaded and that Glacier can now assemble the archive from the
// uploaded parts. After assembling and saving the archive to the vault, Glacier
// returns the URI path of the newly created archive resource. Using the URI path,
// you can then access the archive. After you upload an archive, you should save
// the archive ID returned to retrieve the archive at a later point. You can also
// get the vault inventory to obtain a list of archive IDs in a vault. For more
// information, see InitiateJob.
//
// In the request, you must include the computed SHA256 tree hash of the entire
// archive you have uploaded. For information about computing a SHA256 tree hash,
// see [Computing Checksums]. On the server side, Glacier also constructs the SHA256 tree hash of the
// assembled archive. If the values match, Glacier saves the archive to the vault;
// otherwise, it returns an error, and the operation fails. The ListPartsoperation returns
// a list of parts uploaded for a specific multipart upload. It includes checksum
// information for each uploaded part that can be used to debug a bad checksum
// issue.
//
// Additionally, Glacier also checks for any missing content ranges when
// assembling the archive, if missing content ranges are found, Glacier returns an
// error and the operation fails.
//
// Complete Multipart Upload is an idempotent operation. After your first
// successful complete multipart upload, if you call the operation again within a
// short period, the operation will succeed and return the same archive ID. This is
// useful in the event you experience a network issue that causes an aborted
// connection or receive a 500 server error, in which case you can repeat your
// Complete Multipart Upload request and get the same archive ID without creating
// duplicate archives. Note, however, that after the multipart upload completes,
// you cannot call the List Parts operation and the multipart upload will not
// appear in List Multipart Uploads response, even if idempotent complete is
// possible.
//
// An AWS account has full permission to perform all operations (actions).
// However, AWS Identity and Access Management (IAM) users don't have any
// permissions by default. You must grant them explicit permission to perform
// specific actions. For more information, see [Access Control Using AWS Identity and Access Management (IAM)].
//
// For conceptual information and underlying REST API, see [Uploading Large Archives in Parts (Multipart Upload)] and [Complete Multipart Upload] in the Amazon
// Glacier Developer Guide.
//
// [Complete Multipart Upload]: https://docs.aws.amazon.com/amazonglacier/latest/dev/api-multipart-complete-upload.html
// [Uploading Large Archives in Parts (Multipart Upload)]: https://docs.aws.amazon.com/amazonglacier/latest/dev/uploading-archive-mpu.html
// [Access Control Using AWS Identity and Access Management (IAM)]: https://docs.aws.amazon.com/amazonglacier/latest/dev/using-iam-with-amazon-glacier.html
// [Computing Checksums]: https://docs.aws.amazon.com/amazonglacier/latest/dev/checksum-calculations.html
