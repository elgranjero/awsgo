package kms

// GetPublicKey is generated as a reference stub.
// Executable command wiring lives under cmd/kms.go.
//
// Returns the public key of an asymmetric KMS key. Unlike the private key of a
// asymmetric KMS key, which never leaves KMS unencrypted, callers with
// kms:GetPublicKey permission can download the public key of an asymmetric KMS
// key. You can share the public key to allow others to encrypt messages and verify
// signatures outside of KMS. For information about asymmetric KMS keys, see [Asymmetric KMS keys]in
// the Key Management Service Developer Guide.
//
// You do not need to download the public key. Instead, you can use the public key
// within KMS by calling the Encrypt, ReEncrypt, or Verify operations with the identifier of an
// asymmetric KMS key. When you use the public key within KMS, you benefit from the
// authentication, authorization, and logging that are part of every KMS operation.
// You also reduce of risk of encrypting data that cannot be decrypted. These
// features are not effective outside of KMS.
//
// To help you use the public key safely outside of KMS, GetPublicKey returns
// important information about the public key in the response, including:
//
// [KeySpec]
// - : The type of key material in the public key, such as RSA_4096 or
// ECC_NIST_P521 .
//
// [KeyUsage]
// - : Whether the key is used for encryption, signing, or deriving a shared
// secret.
//
// [EncryptionAlgorithms]
// - , [KeyAgreementAlgorithms], or [SigningAlgorithms]: A list of the encryption algorithms, key agreement algorithms, or
// signing algorithms for the key.
//
// Although KMS cannot enforce these restrictions on external operations, it is
// crucial that you use this information to prevent the public key from being used
// improperly. For example, you can prevent a public signing key from being used
// encrypt data, or prevent a public key from being used with an encryption
// algorithm that is not supported by KMS. You can also avoid errors, such as using
// the wrong signing algorithm in a verification operation.
//
// To verify a signature outside of KMS with an SM2 public key (China Regions
// only), you must specify the distinguishing ID. By default, KMS uses
// 1234567812345678 as the distinguishing ID. For more information, see [Offline verification with SM2 key pairs].
//
// The KMS key that you use for this operation must be in a compatible key state.
// For details, see [Key states of KMS keys]in the Key Management Service Developer Guide.
//
// Cross-account use: Yes. To perform this operation with a KMS key in a different
// Amazon Web Services account, specify the key ARN or alias ARN in the value of
// the KeyId parameter.
//
// Required permissions: [kms:GetPublicKey] (key policy)
//
// Related operations: CreateKey
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [SigningAlgorithms]: https://docs.aws.amazon.com/kms/latest/APIReference/API_GetPublicKey.html#KMS-GetPublicKey-response-SigningAlgorithms
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [kms:GetPublicKey]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [EncryptionAlgorithms]: https://docs.aws.amazon.com/kms/latest/APIReference/API_GetPublicKey.html#KMS-GetPublicKey-response-EncryptionAlgorithms
// [Asymmetric KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/symmetric-asymmetric.html
// [KeyAgreementAlgorithms]: https://docs.aws.amazon.com/kms/latest/APIReference/API_GetPublicKey.html#KMS-GetPublicKey-response-KeyAgreementAlgorithms
// [KeySpec]: https://docs.aws.amazon.com/kms/latest/APIReference/API_GetPublicKey.html#KMS-GetPublicKey-response-KeySpec
// [Offline verification with SM2 key pairs]: https://docs.aws.amazon.com/kms/latest/developerguide/offline-operations.html#key-spec-sm-offline-verification
// [KeyUsage]: https://docs.aws.amazon.com/kms/latest/APIReference/API_GetPublicKey.html#KMS-GetPublicKey-response-KeyUsage
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
