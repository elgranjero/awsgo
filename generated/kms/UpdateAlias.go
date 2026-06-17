package kms

// UpdateAlias is generated as a reference stub.
// Executable command wiring lives under cmd/kms.go.
//
// Associates an existing KMS alias with a different KMS key. Each alias is
// associated with only one KMS key at a time, although a KMS key can have multiple
// aliases. The alias and the KMS key must be in the same Amazon Web Services
// account and Region.
//
// Adding, deleting, or updating an alias can allow or deny permission to the KMS
// key. For details, see [ABAC for KMS]in the Key Management Service Developer Guide.
//
// The current and new KMS key must be the same type (both symmetric or both
// asymmetric or both HMAC), and they must have the same key usage. This
// restriction prevents errors in code that uses aliases. If you must assign an
// alias to a different type of KMS key, use DeleteAliasto delete the old alias and CreateAlias to
// create a new alias.
//
// You cannot use UpdateAlias to change an alias name. To change an alias name,
// use DeleteAliasto delete the old alias and CreateAlias to create a new alias.
//
// Because an alias is not a property of a KMS key, you can create, update, and
// delete the aliases of a KMS key without affecting the KMS key. Also, aliases do
// not appear in the response from the DescribeKeyoperation. To get the aliases of all KMS
// keys in the account, use the ListAliasesoperation.
//
// The KMS key that you use for this operation must be in a compatible key state.
// For details, see [Key states of KMS keys]in the Key Management Service Developer Guide.
//
// Cross-account use: No. You cannot perform this operation on a KMS key in a
// different Amazon Web Services account.
//
// # Required permissions
//
// [kms:UpdateAlias]
// - on the alias (IAM policy).
//
// [kms:UpdateAlias]
// - on the current KMS key (key policy).
//
// [kms:UpdateAlias]
// - on the new KMS key (key policy).
//
// For details, see [Controlling access to aliases] in the Key Management Service Developer Guide.
//
// Related operations:
//
// # CreateAlias
//
// # DeleteAlias
//
// # ListAliases
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [ABAC for KMS]: https://docs.aws.amazon.com/kms/latest/developerguide/abac.html
// [kms:UpdateAlias]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
// [Controlling access to aliases]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-alias.html#alias-access
