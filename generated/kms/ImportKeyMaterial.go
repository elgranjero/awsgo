package kms

// ImportKeyMaterial is generated as a reference stub.
// Executable command wiring lives under cmd/kms.go.
//
// Imports or reimports key material into an existing KMS key that was created
// without key material. You can also use this operation to set or update the
// expiration model and expiration date of the imported key material.
//
// By default, KMS creates KMS keys with key material that it generates. You can
// also generate and import your own key material. For more information about
// importing key material, see [Importing key material].
//
// For asymmetric and HMAC keys, you cannot change the key material after the
// initial import. You can import multiple key materials into symmetric encryption
// keys and rotate the key material on demand using RotateKeyOnDemand .
//
// You can import new key materials into multi-Region symmetric encryption keys.
// To do so, you must import the new key material into the primary Region key. Then
// you can import the same key materials into the replica Region keys. You cannot
// directly import new key material into the replica Region keys.
//
// To import new key material for a multi-Region symmetric key, you’ll need to
// complete the following:
//
// - Call ImportKeyMaterial on the primary Region key with the ImportType set to
// NEW_KEY_MATERIAL .
//
// - Call ImportKeyMaterial on the replica Region key with the ImportType set to
// EXISTING_KEY_MATERIAL using the same key material imported to the primary
// Region key. You must do this for every replica Region key before you can perform
// the RotateKeyOnDemandoperation on the primary Region key.
//
// After you import key material, you can [reimport the same key material] into that KMS key or, if the key
// supports on-demand rotation, import new key material. You can use the ImportType
// parameter to indicate whether you are importing new key material or re-importing
// previously imported key material. You might reimport key material to replace key
// material that expired or key material that you deleted. You might also reimport
// key material to change the expiration model or expiration date of the key
// material.
//
// Each time you import key material into KMS, you can determine whether (
// ExpirationModel ) and when ( ValidTo ) the key material expires. To change the
// expiration of your key material, you must import it again, either by calling
// ImportKeyMaterial or using the [import features] of the KMS console.
//
// Before you call ImportKeyMaterial , complete these steps:
//
// - Create or identify a KMS key with EXTERNAL origin, which indicates that the
// KMS key is designed for imported key material.
//
// To create a new KMS key for imported key material, call the CreateKeyoperation with an
//
// Origin value of EXTERNAL . You can create a symmetric encryption KMS key, HMAC
// KMS key, asymmetric encryption KMS key, asymmetric key agreement key, or
// asymmetric signing KMS key. You can also import key material into a [multi-Region key]of any
// supported type. However, you can't import key material into a KMS key in a [custom key store].
//
// - Call the GetParametersForImportoperation to get a public key and import token set for importing
// key material.
//
// - Use the public key in the GetParametersForImportresponse to encrypt your key material.
//
// Then, in an ImportKeyMaterial request, you submit your encrypted key material
// and import token. When calling this operation, you must specify the following
// values:
//
// - The key ID or key ARN of the KMS key to associate with the imported key
// material. Its Origin must be EXTERNAL and its KeyState must be PendingImport
// or Enabled . You cannot perform this operation on a KMS key in a [custom key store], or on a
// KMS key in a different Amazon Web Services account. To get the Origin and
// KeyState of a KMS key, call DescribeKey.
//
// - The encrypted key material.
//
// - The import token that GetParametersForImportreturned. You must use a public key and token from
// the same GetParametersForImport response.
//
// - Whether the key material expires ( ExpirationModel ) and, if so, when (
// ValidTo ). For help with this choice, see [Setting an expiration time]in the Key Management Service
// Developer Guide.
//
// If you set an expiration date, KMS deletes the key material from the KMS key on
//
// the specified date, making the KMS key unusable. To use the KMS key in
// cryptographic operations again, you must reimport the same key material.
// However, you can delete and reimport the key material at any time, including
// before the key material expires. Each time you reimport, you can eliminate or
// reset the expiration time.
//
// When this operation is successful, the state of the KMS key changes to Enabled ,
// and you can use the KMS key in cryptographic operations. For symmetric
// encryption keys, you will need to import all of the key materials associated
// with the KMS key to change its state to Enabled . Use the ListKeyRotations
// operation to list the ID and import state of each key material associated with a
// KMS key.
//
// If this operation fails, use the exception to help determine the problem. If
// the error is related to the key material, the import token, or wrapping key, use
// GetParametersForImportto get a new public key and import token for the KMS key and repeat the import
// procedure. For help, see [Create a KMS key with imported key material]in the Key Management Service Developer Guide.
//
// The KMS key that you use for this operation must be in a compatible key state.
// For details, see [Key states of KMS keys]in the Key Management Service Developer Guide.
//
// Cross-account use: No. You cannot perform this operation on a KMS key in a
// different Amazon Web Services account.
//
// Required permissions: [kms:ImportKeyMaterial] (key policy)
//
// Related operations:
//
// # DeleteImportedKeyMaterial
//
// # GetParametersForImport
//
// # ListKeyRotations
//
// # RotateKeyOnDemand
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [Importing key material]: https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys.html
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [kms:ImportKeyMaterial]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [reimport the same key material]: https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys-import-key-material.html#reimport-key-material
// [import features]: https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys-import-key-material.html#importing-keys-import-key-material-console
// [Create a KMS key with imported key material]: https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys-conceptual.html
// [Setting an expiration time]: https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys-import-key-material.html#importing-keys-expiration
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
// [custom key store]: https://docs.aws.amazon.com/kms/latest/developerguide/key-store-overview.html
//
// [multi-Region key]: https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html
