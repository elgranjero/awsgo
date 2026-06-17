package kms

// RotateKeyOnDemand is generated as a reference stub.
// Executable command wiring lives under cmd/kms.go.
//
// Immediately initiates rotation of the key material of the specified symmetric
// encryption KMS key.
//
// You can perform [on-demand rotation] of the key material in customer managed KMS keys, regardless
// of whether or not [automatic key rotation]is enabled. On-demand rotations do not change existing
// automatic rotation schedules. For example, consider a KMS key that has automatic
// key rotation enabled with a rotation period of 730 days. If the key is scheduled
// to automatically rotate on April 14, 2024, and you perform an on-demand rotation
// on April 10, 2024, the key will automatically rotate, as scheduled, on April 14,
// 2024 and every 730 days thereafter.
//
// You can perform on-demand key rotation a maximum of 25 times per KMS key. You
// can use the KMS console to view the number of remaining on-demand rotations
// available for a KMS key.
//
// You can use GetKeyRotationStatus to identify any in progress on-demand rotations. You can use ListKeyRotations to
// identify the date that completed on-demand rotations were performed. You can
// monitor rotation of the key material for your KMS keys in CloudTrail and Amazon
// CloudWatch.
//
// On-demand key rotation is supported only on symmetric encryption KMS keys. You
// cannot perform on-demand rotation of [asymmetric KMS keys], [HMAC KMS keys], or KMS keys in a [custom key store]. When you initiate
// on-demand key rotation on a symmetric encryption KMS key with imported key
// material, you must have already imported [new key material]and that key material's state should
// be PENDING_ROTATION . Use the ListKeyRotations operation to check the state of
// all key materials associated with a KMS key. To perform on-demand rotation of a
// set of related [multi-Region keys], import new key material in the primary Region key, import the
// same key material in each replica Region key, and invoke the on-demand rotation
// on the primary Region key.
//
// You cannot initiate on-demand rotation of [Amazon Web Services managed KMS keys]. KMS always rotates the key material
// of Amazon Web Services managed keys every year. Rotation of [Amazon Web Services owned KMS keys]is managed by the
// Amazon Web Services service that owns the key.
//
// The KMS key that you use for this operation must be in a compatible key state.
// For details, see [Key states of KMS keys]in the Key Management Service Developer Guide.
//
// Cross-account use: No. You cannot perform this operation on a KMS key in a
// different Amazon Web Services account.
//
// Required permissions: [kms:RotateKeyOnDemand] (key policy)
//
// Related operations:
//
// # EnableKeyRotation
//
// # DisableKeyRotation
//
// # GetKeyRotationStatus
//
// # ImportKeyMaterial
//
// # ListKeyRotations
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [new key material]: https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys-import-key-material.html
// [HMAC KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/hmac.html
// [Amazon Web Services managed KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-managed-key
// [on-demand rotation]: https://docs.aws.amazon.com/kms/latest/developerguide/rotating-keys-on-demand.html
// [asymmetric KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/symmetric-asymmetric.html
// [Amazon Web Services owned KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-owned-key
// [automatic key rotation]: https://docs.aws.amazon.com/kms/latest/developerguide/rotating-keys-enable-disable.html
// [kms:RotateKeyOnDemand]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [multi-Region keys]: https://docs.aws.amazon.com/kms/latest/developerguide/rotate-keys.html#multi-region-rotate
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
// [custom key store]: https://docs.aws.amazon.com/kms/latest/developerguide/key-store-overview.html
