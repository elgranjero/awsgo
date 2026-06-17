package glacier

// CreateVault is generated as a reference stub.
// Executable command wiring lives under cmd/glacier.go.
//
// This operation creates a new vault with the specified name. The name of the
// vault must be unique within a region for an AWS account. You can create up to
// 1,000 vaults per account. If you need to create more vaults, contact Amazon
// Glacier.
//
// You must use the following guidelines when naming a vault.
//
// - Names can be between 1 and 255 characters long.
//
// - Allowed characters are a-z, A-Z, 0-9, '_' (underscore), '-' (hyphen), and
// '.' (period).
//
// This operation is idempotent.
//
// An AWS account has full permission to perform all operations (actions).
// However, AWS Identity and Access Management (IAM) users don't have any
// permissions by default. You must grant them explicit permission to perform
// specific actions. For more information, see [Access Control Using AWS Identity and Access Management (IAM)].
//
// For conceptual information and underlying REST API, see [Creating a Vault in Amazon Glacier] and [Create Vault] in the Amazon
// Glacier Developer Guide.
//
// [Creating a Vault in Amazon Glacier]: https://docs.aws.amazon.com/amazonglacier/latest/dev/creating-vaults.html
// [Access Control Using AWS Identity and Access Management (IAM)]: https://docs.aws.amazon.com/amazonglacier/latest/dev/using-iam-with-amazon-glacier.html
// [Create Vault]: https://docs.aws.amazon.com/amazonglacier/latest/dev/api-vault-put.html
