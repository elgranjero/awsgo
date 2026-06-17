package kms

// DeleteImportedKeyMaterial is generated as a reference stub.
// Executable command wiring lives under cmd/kms.go.
//
// Deletes key material that was previously imported. This operation makes the
// specified KMS key temporarily unusable. To restore the usability of the KMS key,
// reimport the same key material. For more information about importing key
// material into KMS, see [Importing Key Material]in the Key Management Service Developer Guide.
//
// When the specified KMS key is in the PendingDeletion state, this operation does
// not change the KMS key's state. Otherwise, it changes the KMS key's state to
// PendingImport .
//
// Considerations for multi-Region symmetric encryption keys
//
// - When you delete the key material of a primary Region key that is in
// PENDING_ROTATION or PENDING_MULTI_REGION_IMPORT_AND_ROTATION state, you'll
// also be deleting the key materials for the replica Region keys.
//
// - If you delete any key material of a replica Region key, the primary Region
// key and other replica Region keys remain unchanged.
//
// The KMS key that you use for this operation must be in a compatible key state.
// For details, see [Key states of KMS keys]in the Key Management Service Developer Guide.
//
// Cross-account use: No. You cannot perform this operation on a KMS key in a
// different Amazon Web Services account.
//
// Required permissions: [kms:DeleteImportedKeyMaterial] (key policy)
//
// Related operations:
//
// # GetParametersForImport
//
// # ListKeyRotations
//
// # ImportKeyMaterial
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [kms:DeleteImportedKeyMaterial]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [Importing Key Material]: https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
