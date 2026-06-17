package kms

// DisableKeyRotation is generated as a reference stub.
// Executable command wiring lives under cmd/kms.go.
//
// Disables [automatic rotation of the key material] of the specified symmetric encryption KMS key.
//
// Automatic key rotation is supported only on symmetric encryption KMS keys. You
// cannot enable automatic rotation of [asymmetric KMS keys], [HMAC KMS keys], KMS keys with [imported key material], or KMS keys in a [custom key store]. To
// enable or disable automatic rotation of a set of related [multi-Region keys], set the property on
// the primary key.
//
// You can enable (EnableKeyRotation ) and disable automatic rotation of the key material in [customer managed KMS keys]. Key
// material rotation of [Amazon Web Services managed KMS keys]is not configurable. KMS always rotates the key material
// for every year. Rotation of [Amazon Web Services owned KMS keys]varies.
//
// In May 2022, KMS changed the rotation schedule for Amazon Web Services managed
// keys from every three years to every year. For details, see EnableKeyRotation.
//
// The KMS key that you use for this operation must be in a compatible key state.
// For details, see [Key states of KMS keys]in the Key Management Service Developer Guide.
//
// Cross-account use: No. You cannot perform this operation on a KMS key in a
// different Amazon Web Services account.
//
// Required permissions: [kms:DisableKeyRotation] (key policy)
//
// Related operations:
//
// # EnableKeyRotation
//
// # GetKeyRotationStatus
//
// # ListKeyRotations
//
// # RotateKeyOnDemand
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [imported key material]: https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys.html
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [HMAC KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/hmac.html
// [Amazon Web Services managed KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-managed-key
// [automatic rotation of the key material]: https://docs.aws.amazon.com/kms/latest/developerguide/rotating-keys-enable-disable.html
// [asymmetric KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/symmetric-asymmetric.html
// [customer managed KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#customer-mgn-key
// [Amazon Web Services owned KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-owned-key
// [kms:DisableKeyRotation]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [multi-Region keys]: https://docs.aws.amazon.com/kms/latest/developerguide/rotate-keys.html#multi-region-rotate
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
// [custom key store]: https://docs.aws.amazon.com/kms/latest/developerguide/key-store-overview.html
