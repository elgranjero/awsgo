package kms

// DeleteAlias is generated as a reference stub.
// Executable command wiring lives under cmd/kms.go.
//
// Deletes the specified alias.
//
// Adding, deleting, or updating an alias can allow or deny permission to the KMS
// key. For details, see [ABAC for KMS]in the Key Management Service Developer Guide.
//
// Because an alias is not a property of a KMS key, you can delete and change the
// aliases of a KMS key without affecting the KMS key. Also, aliases do not appear
// in the response from the DescribeKeyoperation. To get the aliases of all KMS keys, use the ListAliases
// operation.
//
// Each KMS key can have multiple aliases. To change the alias of a KMS key, use DeleteAlias
// to delete the current alias and CreateAliasto create a new alias. To associate an existing
// alias with a different KMS key, call UpdateAlias.
//
// Cross-account use: No. You cannot perform this operation on an alias in a
// different Amazon Web Services account.
//
// # Required permissions
//
// [kms:DeleteAlias]
// - on the alias (IAM policy).
//
// [kms:DeleteAlias]
// - on the KMS key (key policy).
//
// For details, see [Controlling access to aliases] in the Key Management Service Developer Guide.
//
// Related operations:
//
// # CreateAlias
//
// # ListAliases
//
// # UpdateAlias
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [ABAC for KMS]: https://docs.aws.amazon.com/kms/latest/developerguide/abac.html
// [kms:DeleteAlias]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
// [Controlling access to aliases]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-alias.html#alias-access
