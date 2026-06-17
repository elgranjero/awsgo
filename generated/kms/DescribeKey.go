package kms

// DescribeKey is generated as a reference stub.
// Executable command wiring lives under cmd/kms.go.
//
// Provides detailed information about a KMS key. You can run DescribeKey on a [customer managed key] or
// an [Amazon Web Services managed key].
//
// This detailed information includes the key ARN, creation date (and deletion
// date, if applicable), the key state, and the origin and expiration date (if any)
// of the key material. It includes fields, like KeySpec , that help you
// distinguish different types of KMS keys. It also displays the key usage
// (encryption, signing, or generating and verifying MACs) and the algorithms that
// the KMS key supports.
//
// For [multi-Region keys], DescribeKey displays the primary key and all related replica keys. For
// KMS keys in [CloudHSM key stores], it includes information about the key store, such as the key
// store ID and the CloudHSM cluster ID. For KMS keys in [external key stores], it includes the custom
// key store ID and the ID of the external key.
//
// DescribeKey does not return the following information:
//
// - Aliases associated with the KMS key. To get this information, use ListAliases.
//
// - Whether automatic key rotation is enabled on the KMS key. To get this
// information, use GetKeyRotationStatus. Also, some key states prevent a KMS key from being
// automatically rotated. For details, see [How key rotation works]in the Key Management Service
// Developer Guide.
//
// - Tags on the KMS key. To get this information, use ListResourceTags.
//
// - Key policies and grants on the KMS key. To get this information, use GetKeyPolicyand ListGrants.
//
// In general, DescribeKey is a non-mutating operation. It returns data about KMS
// keys, but doesn't change them. However, Amazon Web Services services use
// DescribeKey to create [Amazon Web Services managed keys] from a predefined Amazon Web Services alias with no key
// ID.
//
// Cross-account use: Yes. To perform this operation with a KMS key in a different
// Amazon Web Services account, specify the key ARN or alias ARN in the value of
// the KeyId parameter.
//
// Required permissions: [kms:DescribeKey] (key policy)
//
// Related operations:
//
// # GetKeyPolicy
//
// # GetKeyRotationStatus
//
// # ListAliases
//
// # ListGrants
//
// # ListKeys
//
// # ListResourceTags
//
// # ListRetirableGrants
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [CloudHSM key stores]: https://docs.aws.amazon.com/kms/latest/developerguide/keystore-cloudhsm.html
// [external key stores]: https://docs.aws.amazon.com/kms/latest/developerguide/keystore-external.html
// [How key rotation works]: https://docs.aws.amazon.com/kms/latest/developerguide/rotate-keys.html#rotate-keys-how-it-works
// [customer managed key]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#customer-mgn-key
// [kms:DescribeKey]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [multi-Region keys]: https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html
// [Amazon Web Services managed keys]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-managed-key
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
// [Amazon Web Services managed key]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-managed-key
