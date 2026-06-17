package glacier

// DescribeVault is generated as a reference stub.
// Executable command wiring lives under cmd/glacier.go.
//
// This operation returns information about a vault, including the vault's Amazon
// Resource Name (ARN), the date the vault was created, the number of archives it
// contains, and the total size of all the archives in the vault. The number of
// archives and their total size are as of the last inventory generation. This
// means that if you add or remove an archive from a vault, and then immediately
// use Describe Vault, the change in contents will not be immediately reflected. If
// you want to retrieve the latest inventory of the vault, use InitiateJob. Amazon Glacier
// generates vault inventories approximately daily. For more information, see [Downloading a Vault Inventory in Amazon Glacier].
//
// An AWS account has full permission to perform all operations (actions).
// However, AWS Identity and Access Management (IAM) users don't have any
// permissions by default. You must grant them explicit permission to perform
// specific actions. For more information, see [Access Control Using AWS Identity and Access Management (IAM)].
//
// For conceptual information and underlying REST API, see [Retrieving Vault Metadata in Amazon Glacier] and [Describe Vault] in the Amazon
// Glacier Developer Guide.
//
// [Retrieving Vault Metadata in Amazon Glacier]: https://docs.aws.amazon.com/amazonglacier/latest/dev/retrieving-vault-info.html
// [Describe Vault]: https://docs.aws.amazon.com/amazonglacier/latest/dev/api-vault-get.html
// [Downloading a Vault Inventory in Amazon Glacier]: https://docs.aws.amazon.com/amazonglacier/latest/dev/vault-inventory.html
// [Access Control Using AWS Identity and Access Management (IAM)]: https://docs.aws.amazon.com/amazonglacier/latest/dev/using-iam-with-amazon-glacier.html
