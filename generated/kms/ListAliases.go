package kms

// ListAliases is generated as a reference stub.
// Executable command wiring lives under cmd/kms.go.
//
// Gets a list of aliases in the caller's Amazon Web Services account and region.
// For more information about aliases, see CreateAlias.
//
// By default, the ListAliases operation returns all aliases in the account and
// region. To get only the aliases associated with a particular KMS key, use the
// KeyId parameter.
//
// The ListAliases response can include aliases that you created and associated
// with your customer managed keys, and aliases that Amazon Web Services created
// and associated with Amazon Web Services managed keys in your account. You can
// recognize Amazon Web Services aliases because their names have the format aws/ ,
// such as aws/dynamodb .
//
// The response might also include aliases that have no TargetKeyId field. These
// are predefined aliases that Amazon Web Services has created but has not yet
// associated with a KMS key. Aliases that Amazon Web Services creates in your
// account, including predefined aliases, do not count against your [KMS aliases quota].
//
// Cross-account use: No. ListAliases does not return aliases in other Amazon Web
// Services accounts.
//
// Required permissions: [kms:ListAliases] (IAM policy)
//
// For details, see [Controlling access to aliases] in the Key Management Service Developer Guide.
//
// Related operations:
//
// # CreateAlias
//
// # DeleteAlias
//
// # UpdateAlias
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [KMS aliases quota]: https://docs.aws.amazon.com/kms/latest/developerguide/resource-limits.html#aliases-per-key
// [kms:ListAliases]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
// [Controlling access to aliases]: https://docs.aws.amazon.com/kms/latest/developerguide/alias-access.html
