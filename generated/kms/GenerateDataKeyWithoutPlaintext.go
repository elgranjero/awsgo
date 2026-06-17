package kms

// GenerateDataKeyWithoutPlaintext is generated as a reference stub.
// Executable command wiring lives under cmd/kms.go.
//
// Returns a unique symmetric data key for use outside of KMS. This operation
// returns a data key that is encrypted under a symmetric encryption KMS key that
// you specify. The bytes in the key are random; they are not related to the caller
// or to the KMS key.
//
// GenerateDataKeyWithoutPlaintext is identical to the GenerateDataKey operation except that it
// does not return a plaintext copy of the data key.
//
// This operation is useful for systems that need to encrypt data at some point,
// but not immediately. When you need to encrypt the data, you call the Decryptoperation
// on the encrypted copy of the key.
//
// It's also useful in distributed systems with different levels of trust. For
// example, you might store encrypted data in containers. One component of your
// system creates new containers and stores an encrypted data key with each
// container. Then, a different component puts the data into the containers. That
// component first decrypts the data key, uses the plaintext data key to encrypt
// data, puts the encrypted data into the container, and then destroys the
// plaintext data key. In this system, the component that creates the containers
// never sees the plaintext data key.
//
// To request an asymmetric data key pair, use the GenerateDataKeyPair or GenerateDataKeyPairWithoutPlaintext operations.
//
// To generate a data key, you must specify the symmetric encryption KMS key that
// is used to encrypt the data key. You cannot use an asymmetric KMS key or a key
// in a custom key store to generate a data key. To get the type of your KMS key,
// use the DescribeKeyoperation.
//
// You must also specify the length of the data key. Use either the KeySpec or
// NumberOfBytes parameters (but not both). For 128-bit and 256-bit data keys, use
// the KeySpec parameter.
//
// To generate an SM4 data key (China Regions only), specify a KeySpec value of
// AES_128 or NumberOfBytes value of 16 . The symmetric encryption key used in
// China Regions to encrypt your data key is an SM4 encryption key.
//
// If the operation succeeds, you will find the encrypted copy of the data key in
// the CiphertextBlob field.
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
// Required permissions: [kms:GenerateDataKeyWithoutPlaintext] (key policy)
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
// # GenerateDataKeyPairWithoutPlaintext
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [Encryption Context]: https://docs.aws.amazon.com/kms/latest/developerguide/encrypt_context.html
// [kms:GenerateDataKeyWithoutPlaintext]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
