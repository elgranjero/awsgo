package kms

// GenerateDataKeyPairWithoutPlaintext is generated as a reference stub.
// Executable command wiring lives under cmd/kms.go.
//
// Returns a unique asymmetric data key pair for use outside of KMS. This
// operation returns a plaintext public key and a copy of the private key that is
// encrypted under the symmetric encryption KMS key you specify. Unlike GenerateDataKeyPair, this
// operation does not return a plaintext private key. The bytes in the keys are
// random; they are not related to the caller or to the KMS key that is used to
// encrypt the private key.
//
// You can use the public key that GenerateDataKeyPairWithoutPlaintext returns to
// encrypt data or verify a signature outside of KMS. Then, store the encrypted
// private key with the data. When you are ready to decrypt data or sign a message,
// you can use the Decryptoperation to decrypt the encrypted private key.
//
// To generate a data key pair, you must specify a symmetric encryption KMS key to
// encrypt the private key in a data key pair. You cannot use an asymmetric KMS key
// or a KMS key in a custom key store. To get the type and origin of your KMS key,
// use the DescribeKeyoperation.
//
// Use the KeyPairSpec parameter to choose an RSA or Elliptic Curve (ECC) data key
// pair. In China Regions, you can also choose an SM2 data key pair. KMS recommends
// that you use ECC key pairs for signing, and use RSA and SM2 key pairs for either
// encryption or signing, but not both. However, KMS cannot enforce any
// restrictions on the use of data key pairs outside of KMS.
//
// GenerateDataKeyPairWithoutPlaintext returns a unique data key pair for each
// request. The bytes in the key are not related to the caller or KMS key that is
// used to encrypt the private key. The public key is a DER-encoded X.509
// SubjectPublicKeyInfo, as specified in [RFC 5280].
//
// You can use an optional encryption context to add additional security to the
// encryption operation. If you specify an EncryptionContext , you must specify the
// same encryption context (a case-sensitive exact match) when decrypting the
// encrypted data key. Otherwise, the request to decrypt fails with an
// InvalidCiphertextException . For more information, see [Encryption Context] in the Key Management
// Service Developer Guide.
//
// The KMS key that you use for this operation must be in a compatible key state.
// For details, see [Key states of KMS keys]in the Key Management Service Developer Guide.
//
// Cross-account use: Yes. To perform this operation with a KMS key in a different
// Amazon Web Services account, specify the key ARN or alias ARN in the value of
// the KeyId parameter.
//
// Required permissions: [kms:GenerateDataKeyPairWithoutPlaintext] (key policy)
//
// Related operations:
//
// # Decrypt
//
// # Encrypt
//
// # GenerateDataKey
//
// # GenerateDataKeyPair
//
// # GenerateDataKeyWithoutPlaintext
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [RFC 5280]: https://tools.ietf.org/html/rfc5280
// [Encryption Context]: https://docs.aws.amazon.com/kms/latest/developerguide/encrypt_context.html
// [kms:GenerateDataKeyPairWithoutPlaintext]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
