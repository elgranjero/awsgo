package kms

// ReEncrypt is generated as a reference stub.
// Executable command wiring lives under cmd/kms.go.
//
// Decrypts ciphertext and then reencrypts it entirely within KMS. You can use
// this operation to change the KMS key under which data is encrypted, such as when
// you [manually rotate]a KMS key or change the KMS key that protects a ciphertext. You can also
// use it to reencrypt ciphertext under the same KMS key, such as to change the [encryption context]of
// a ciphertext.
//
// The ReEncrypt operation can decrypt ciphertext that was encrypted by using a
// KMS key in an KMS operation, such as Encryptor GenerateDataKey. It can also decrypt ciphertext that
// was encrypted by using the public key of an [asymmetric KMS key]outside of KMS. However, it cannot
// decrypt ciphertext produced by other libraries, such as the [Amazon Web Services Encryption SDK]or [Amazon S3 client-side encryption]. These
// libraries return a ciphertext format that is incompatible with KMS.
//
// When you use the ReEncrypt operation, you need to provide information for the
// decrypt operation and the subsequent encrypt operation.
//
// - If your ciphertext was encrypted under an asymmetric KMS key, you must use
// the SourceKeyId parameter to identify the KMS key that encrypted the
// ciphertext. You must also supply the encryption algorithm that was used. This
// information is required to decrypt the data.
//
// - If your ciphertext was encrypted under a symmetric encryption KMS key, the
// SourceKeyId parameter is optional. KMS can get this information from metadata
// that it adds to the symmetric ciphertext blob. This feature adds durability to
// your implementation by ensuring that authorized users can decrypt ciphertext
// decades after it was encrypted, even if they've lost track of the key ID.
// However, specifying the source KMS key is always recommended as a best practice.
// When you use the SourceKeyId parameter to specify a KMS key, KMS uses only the
// KMS key you specify. If the ciphertext was encrypted under a different KMS key,
// the ReEncrypt operation fails. This practice ensures that you use the KMS key
// that you intend.
//
// - To reencrypt the data, you must use the DestinationKeyId parameter to
// specify the KMS key that re-encrypts the data after it is decrypted. If the
// destination KMS key is an asymmetric KMS key, you must also provide the
// encryption algorithm. The algorithm that you choose must be compatible with the
// KMS key.
//
// When you use an asymmetric KMS key to encrypt or reencrypt data, be sure to
//
// record the KMS key and encryption algorithm that you choose. You will be
// required to provide the same KMS key and encryption algorithm when you decrypt
// the data. If the KMS key and algorithm do not match the values used to encrypt
// the data, the decrypt operation fails.
//
// You are not required to supply the key ID and encryption algorithm when you
//
// decrypt with symmetric encryption KMS keys because KMS stores this information
// in the ciphertext blob. KMS cannot store metadata in ciphertext generated with
// asymmetric keys. The standard format for asymmetric key ciphertext does not
// include configurable fields.
//
// The KMS key that you use for this operation must be in a compatible key state.
// For details, see [Key states of KMS keys]in the Key Management Service Developer Guide.
//
// Cross-account use: Yes. The source KMS key and destination KMS key can be in
// different Amazon Web Services accounts. Either or both KMS keys can be in a
// different account than the caller. To specify a KMS key in a different account,
// you must use its key ARN or alias ARN.
//
// Required permissions:
//
// [kms:ReEncryptFrom]
// - permission on the source KMS key (key policy)
//
// [kms:ReEncryptTo]
// - permission on the destination KMS key (key policy)
//
// To permit reencryption from or to a KMS key, include the "kms:ReEncrypt*"
// permission in your [key policy]. This permission is automatically included in the key
// policy when you use the console to create a KMS key. But you must include it
// manually when you create a KMS key programmatically or when you use the PutKeyPolicy
// operation to set a key policy.
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
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [Amazon Web Services Encryption SDK]: https://docs.aws.amazon.com/encryption-sdk/latest/developer-guide/
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [asymmetric KMS key]: https://docs.aws.amazon.com/kms/latest/developerguide/symmetric-asymmetric.html
// [key policy]: https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html
// [Amazon S3 client-side encryption]: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingClientSideEncryption.html
// [kms:ReEncryptTo]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [encryption context]: https://docs.aws.amazon.com/kms/latest/developerguide/encrypt_context.html
// [manually rotate]: https://docs.aws.amazon.com/kms/latest/developerguide/rotate-keys-manually.html
// [kms:ReEncryptFrom]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
