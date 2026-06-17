package kms

// CreateAlias is generated as a reference stub.
// Executable command wiring lives under cmd/kms.go.
//
// Creates a friendly name for a KMS key.
//
// Adding, deleting, or updating an alias can allow or deny permission to the KMS
// key. For details, see [ABAC for KMS]in the Key Management Service Developer Guide.
//
// You can use an alias to identify a KMS key in the KMS console, in the DescribeKey
// operation and in [cryptographic operations], such as Encrypt and GenerateDataKey. You can also change the KMS key that's
// associated with the alias (UpdateAlias ) or delete the alias (DeleteAlias ) at any time. These
// operations don't affect the underlying KMS key.
//
// You can associate the alias with any customer managed key in the same Amazon
// Web Services Region. Each alias is associated with only one KMS key at a time,
// but a KMS key can have multiple aliases. A valid KMS key is required. You can't
// create an alias without a KMS key.
//
// The alias must be unique in the account and Region, but you can have aliases
// with the same name in different Regions. For detailed information about aliases,
// see [Aliases in KMS]in the Key Management Service Developer Guide.
//
// This operation does not return a response. To get the alias that you created,
// use the ListAliasesoperation.
//
// The KMS key that you use for this operation must be in a compatible key state.
// For details, see [Key states of KMS keys]in the Key Management Service Developer Guide.
//
// Cross-account use: No. You cannot perform this operation on an alias in a
// different Amazon Web Services account.
//
// # Required permissions
//
// [kms:CreateAlias]
// - on the alias (IAM policy).
//
// [kms:CreateAlias]
// - on the KMS key (key policy).
//
// For details, see [Controlling access to aliases] in the Key Management Service Developer Guide.
//
// Related operations:
//
// # DeleteAlias
//
// # ListAliases
//
// # UpdateAlias
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [cryptographic operations]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-cryptography.html#cryptographic-operations
// [kms:CreateAlias]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [Aliases in KMS]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-alias.html
// [ABAC for KMS]: https://docs.aws.amazon.com/kms/latest/developerguide/abac.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
// [Controlling access to aliases]: https://docs.aws.amazon.com/kms/latest/developerguide/alias-access.html
