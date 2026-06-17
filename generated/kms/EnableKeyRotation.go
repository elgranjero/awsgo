package kms

// EnableKeyRotation is generated as a reference stub.
// Executable command wiring lives under cmd/kms.go.
//
// Enables [automatic rotation of the key material] of the specified symmetric encryption KMS key.
//
// By default, when you enable automatic rotation of a [customer managed KMS key], KMS rotates the key
// material of the KMS key one year (approximately 365 days) from the enable date
// and every year thereafter. You can use the optional RotationPeriodInDays
// parameter to specify a custom rotation period when you enable key rotation, or
// you can use RotationPeriodInDays to modify the rotation period of a key that
// you previously enabled automatic key rotation on.
//
// You can monitor rotation of the key material for your KMS keys in CloudTrail
// and Amazon CloudWatch. To disable rotation of the key material in a customer
// managed KMS key, use the DisableKeyRotationoperation. You can use the GetKeyRotationStatus operation to identify any
// in progress rotations. You can use the ListKeyRotationsoperation to view the details of
// completed rotations.
//
// Automatic key rotation is supported only on symmetric encryption KMS keys. You
// cannot enable automatic rotation of [asymmetric KMS keys], [HMAC KMS keys], KMS keys with [imported key material], or KMS keys in a [custom key store]. To
// enable or disable automatic rotation of a set of related [multi-Region keys], set the property on
// the primary key.
//
// You cannot enable or disable automatic rotation of [Amazon Web Services managed KMS keys]. KMS always rotates the key
// material of Amazon Web Services managed keys every year. Rotation of [Amazon Web Services owned KMS keys]is managed
// by the Amazon Web Services service that owns the key.
//
// In May 2022, KMS changed the rotation schedule for Amazon Web Services managed
// keys from every three years (approximately 1,095 days) to every year
// (approximately 365 days).
//
// New Amazon Web Services managed keys are automatically rotated one year after
// they are created, and approximately every year thereafter.
//
// Existing Amazon Web Services managed keys are automatically rotated one year
// after their most recent rotation, and every year thereafter.
//
// The KMS key that you use for this operation must be in a compatible key state.
// For details, see [Key states of KMS keys]in the Key Management Service Developer Guide.
//
// Cross-account use: No. You cannot perform this operation on a KMS key in a
// different Amazon Web Services account.
//
// Required permissions: [kms:EnableKeyRotation] (key policy)
//
// Related operations:
//
// # DisableKeyRotation
//
// # GetKeyRotationStatus
//
// # ListKeyRotations
//
// RotateKeyOnDemand
//
// - You can perform on-demand (RotateKeyOnDemand ) rotation of the key material in customer
// managed KMS keys, regardless of whether or not automatic key rotation is
// enabled.
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [kms:EnableKeyRotation]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [imported key material]: https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys.html
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [HMAC KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/hmac.html
// [Amazon Web Services managed KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-managed-key
// [customer managed KMS key]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#customer-mgn-key
// [automatic rotation of the key material]: https://docs.aws.amazon.com/kms/latest/developerguide/rotating-keys-enable-disable.html
// [asymmetric KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/symmetric-asymmetric.html
// [Amazon Web Services owned KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-owned-key
// [multi-Region keys]: https://docs.aws.amazon.com/kms/latest/developerguide/rotate-keys.html#multi-region-rotate
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
// [custom key store]: https://docs.aws.amazon.com/kms/latest/developerguide/key-store-overview.html
