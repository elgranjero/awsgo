package kms

// Encrypt is generated as a reference stub.
// Executable command wiring lives under cmd/kms.go.
//
// Encrypts plaintext of up to 4,096 bytes using a KMS key. You can use a
// symmetric or asymmetric KMS key with a KeyUsage of ENCRYPT_DECRYPT .
//
// You can use this operation to encrypt small amounts of arbitrary data, such as
// a personal identifier or database password, or other sensitive information. You
// don't need to use the Encrypt operation to encrypt a data key. The GenerateDataKey and GenerateDataKeyPair
// operations return a plaintext data key and an encrypted copy of that data key.
//
// If you use a symmetric encryption KMS key, you can use an encryption context to
// add additional security to your encryption operation. If you specify an
// EncryptionContext when encrypting data, you must specify the same encryption
// context (a case-sensitive exact match) when decrypting the data. Otherwise, the
// request to decrypt fails with an InvalidCiphertextException . For more
// information, see [Encryption Context]in the Key Management Service Developer Guide.
//
// If you specify an asymmetric KMS key, you must also specify the encryption
// algorithm. The algorithm must be compatible with the KMS key spec.
//
// When you use an asymmetric KMS key to encrypt or reencrypt data, be sure to
// record the KMS key and encryption algorithm that you choose. You will be
// required to provide the same KMS key and encryption algorithm when you decrypt
// the data. If the KMS key and algorithm do not match the values used to encrypt
// the data, the decrypt operation fails.
//
// You are not required to supply the key ID and encryption algorithm when you
// decrypt with symmetric encryption KMS keys because KMS stores this information
// in the ciphertext blob. KMS cannot store metadata in ciphertext generated with
// asymmetric keys. The standard format for asymmetric key ciphertext does not
// include configurable fields.
//
// The maximum size of the data that you can encrypt varies with the type of KMS
// key and the encryption algorithm that you choose.
//
// - Symmetric encryption KMS keys
//
// - SYMMETRIC_DEFAULT : 4096 bytes
//
// - RSA_2048
//
// - RSAES_OAEP_SHA_1 : 214 bytes
//
// - RSAES_OAEP_SHA_256 : 190 bytes
//
// - RSA_3072
//
// - RSAES_OAEP_SHA_1 : 342 bytes
//
// - RSAES_OAEP_SHA_256 : 318 bytes
//
// - RSA_4096
//
// - RSAES_OAEP_SHA_1 : 470 bytes
//
// - RSAES_OAEP_SHA_256 : 446 bytes
//
// - SM2PKE : 1024 bytes (China Regions only)
//
// The KMS key that you use for this operation must be in a compatible key state.
// For details, see [Key states of KMS keys]in the Key Management Service Developer Guide.
//
// Cross-account use: Yes. To perform this operation with a KMS key in a different
// Amazon Web Services account, specify the key ARN or alias ARN in the value of
// the KeyId parameter.
//
// Required permissions: [kms:Encrypt] (key policy)
//
// Related operations:
//
// # Decrypt
//
// # GenerateDataKey
//
// # GenerateDataKeyPair
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [Encryption Context]: https://docs.aws.amazon.com/kms/latest/developerguide/encrypt_context.html
// [kms:Encrypt]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
