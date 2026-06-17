package kms

// GetParametersForImport is generated as a reference stub.
// Executable command wiring lives under cmd/kms.go.
//
// Returns the public key and an import token you need to import or reimport key
// material for a KMS key.
//
// By default, KMS keys are created with key material that KMS generates. This
// operation supports [Importing key material], an advanced feature that lets you generate and import the
// cryptographic key material for a KMS key.
//
// Before calling GetParametersForImport , use the CreateKey operation with an Origin value
// of EXTERNAL to create a KMS key with no key material. You can import key
// material for a symmetric encryption KMS key, HMAC KMS key, asymmetric encryption
// KMS key, or asymmetric signing KMS key. You can also import key material into a [multi-Region key]
// of any supported type. However, you can't import key material into a KMS key in
// a [custom key store]. You can also use GetParametersForImport to get a public key and import
// token to [reimport the original key material]into a KMS key whose key material expired or was deleted.
//
// GetParametersForImport returns the items that you need to import your key
// material.
//
// - The public key (or "wrapping key") of an RSA key pair that KMS generates.
//
// You will use this public key to encrypt ("wrap") your key material while it's
//
// in transit to KMS.
//
// - A import token that ensures that KMS can decrypt your key material and
// associate it with the correct KMS key.
//
// The public key and its import token are permanently linked and must be used
// together. Each public key and import token set is valid for 24 hours. The
// expiration date and time appear in the ParametersValidTo field in the
// GetParametersForImport response. You cannot use an expired public key or import
// token in an ImportKeyMaterialrequest. If your key and token expire, send another
// GetParametersForImport request.
//
// GetParametersForImport requires the following information:
//
// - The key ID of the KMS key for which you are importing the key material.
//
// - The key spec of the public key ("wrapping key") that you will use to
// encrypt your key material during import.
//
// - The wrapping algorithm that you will use with the public key to encrypt
// your key material.
//
// You can use the same or a different public key spec and wrapping algorithm each
// time you import or reimport the same key material.
//
// The KMS key that you use for this operation must be in a compatible key state.
// For details, see [Key states of KMS keys]in the Key Management Service Developer Guide.
//
// Cross-account use: No. You cannot perform this operation on a KMS key in a
// different Amazon Web Services account.
//
// Required permissions: [kms:GetParametersForImport] (key policy)
//
// Related operations:
//
// # ImportKeyMaterial
//
// # DeleteImportedKeyMaterial
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [Importing key material]: https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys.html
// [kms:GetParametersForImport]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [reimport the original key material]: https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys-import-key-material.html#reimport-key-material
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
// [multi-Region key]: https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html
// [custom key store]: https://docs.aws.amazon.com/kms/latest/developerguide/key-store-overview.html
