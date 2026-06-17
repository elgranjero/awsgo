package glacier

// DeleteVault is generated as a reference stub.
// Executable command wiring lives under cmd/glacier.go.
//
// This operation deletes a vault. Amazon Glacier will delete a vault only if
// there are no archives in the vault as of the last inventory and there have been
// no writes to the vault since the last inventory. If either of these conditions
// is not satisfied, the vault deletion fails (that is, the vault is not removed)
// and Amazon Glacier returns an error. You can use DescribeVaultto return the number of
// archives in a vault, and you can use [Initiate a Job (POST jobs)]to initiate a new inventory retrieval for
// a vault. The inventory contains the archive IDs you use to delete archives using
// [Delete Archive (DELETE archive)].
//
// This operation is idempotent.
//
// An AWS account has full permission to perform all operations (actions).
// However, AWS Identity and Access Management (IAM) users don't have any
// permissions by default. You must grant them explicit permission to perform
// specific actions. For more information, see [Access Control Using AWS Identity and Access Management (IAM)].
//
// For conceptual information and underlying REST API, see [Deleting a Vault in Amazon Glacier] and [Delete Vault] in the Amazon
// Glacier Developer Guide.
//
// [Delete Archive (DELETE archive)]: https://docs.aws.amazon.com/amazonglacier/latest/dev/api-archive-delete.html
// [Deleting a Vault in Amazon Glacier]: https://docs.aws.amazon.com/amazonglacier/latest/dev/deleting-vaults.html
// [Delete Vault]: https://docs.aws.amazon.com/amazonglacier/latest/dev/api-vault-delete.html
// [Initiate a Job (POST jobs)]: https://docs.aws.amazon.com/amazonglacier/latest/dev/api-initiate-job-post.html
// [Access Control Using AWS Identity and Access Management (IAM)]: https://docs.aws.amazon.com/amazonglacier/latest/dev/using-iam-with-amazon-glacier.html
