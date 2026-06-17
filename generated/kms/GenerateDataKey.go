package kms

// GenerateDataKey is generated as a reference stub.
// Executable command wiring lives under cmd/kms.go.
//
// Returns a unique symmetric data key for use outside of KMS. This operation
// returns a plaintext copy of the data key and a copy that is encrypted under a
// symmetric encryption KMS key that you specify. The bytes in the plaintext key
// are random; they are not related to the caller or the KMS key. You can use the
// plaintext key to encrypt your data outside of KMS and store the encrypted data
// key with the encrypted data.
//
// To generate a data key, specify the symmetric encryption KMS key that will be
// used to encrypt the data key. You cannot use an asymmetric KMS key to encrypt
// data keys. To get the type of your KMS key, use the DescribeKeyoperation.
//
// You must also specify the length of the data key. Use either the KeySpec or
// NumberOfBytes parameters (but not both). For 128-bit and 256-bit data keys, use
// the KeySpec parameter.
//
// To generate a 128-bit SM4 data key (China Regions only), specify a KeySpec
// value of AES_128 or a NumberOfBytes value of 16 . The symmetric encryption key
// used in China Regions to encrypt your data key is an SM4 encryption key.
//
// To get only an encrypted copy of the data key, use GenerateDataKeyWithoutPlaintext. To generate an asymmetric
// data key pair, use the GenerateDataKeyPairor GenerateDataKeyPairWithoutPlaintext operation. To get a cryptographically secure random
// byte string, use GenerateRandom.
//
// You can use an optional encryption context to add additional security to the
// encryption operation. If you specify an EncryptionContext , you must specify the
// same encryption context (a case-sensitive exact match) when decrypting the
// encrypted data key. Otherwise, the request to decrypt fails with an
// InvalidCiphertextException . For more information, see [Encryption Context] in the Key Management
// Service Developer Guide.
//
// GenerateDataKey also supports [Amazon Web Services Nitro Enclaves], which provide an isolated compute environment
// in Amazon EC2. To call GenerateDataKey for an Amazon Web Services Nitro enclave
// or NitroTPM, use the [Amazon Web Services Nitro Enclaves SDK]or any Amazon Web Services SDK. Use the Recipient
// parameter to provide the attestation document for the attested environment.
// GenerateDataKey returns a copy of the data key encrypted under the specified KMS
// key, as usual. But instead of a plaintext copy of the data key, the response
// includes a copy of the data key encrypted under the public key from the
// attestation document ( CiphertextForRecipient ). For information about the
// interaction between KMS and Amazon Web Services Nitro Enclaves or Amazon Web
// Services NitroTPM, see [Cryptographic attestation support in KMS]in the Key Management Service Developer Guide.
//
// The KMS key that you use for this operation must be in a compatible key state.
// For details, see [Key states of KMS keys]in the Key Management Service Developer Guide.
//
// # How to use your data key
//
// We recommend that you use the following pattern to encrypt data locally in your
// application. You can write your own code or use a client-side encryption
// library, such as the [Amazon Web Services Encryption SDK], the [Amazon DynamoDB Encryption Client], or [Amazon S3 client-side encryption] to do these tasks for you.
//
// To encrypt data outside of KMS:
//
// - Use the GenerateDataKey operation to get a data key.
//
// - Use the plaintext data key (in the Plaintext field of the response) to
// encrypt your data outside of KMS. Then erase the plaintext data key from memory.
//
// - Store the encrypted data key (in the CiphertextBlob field of the response)
// with the encrypted data.
//
// To decrypt data outside of KMS:
//
// - Use the Decryptoperation to decrypt the encrypted data key. The operation returns
// a plaintext copy of the data key.
//
// - Use the plaintext data key to decrypt data outside of KMS, then erase the
// plaintext data key from memory.
//
// Cross-account use: Yes. To perform this operation with a KMS key in a different
// Amazon Web Services account, specify the key ARN or alias ARN in the value of
// the KeyId parameter.
//
// Required permissions: [kms:GenerateDataKey] (key policy)
//
// Related operations:
//
// # Decrypt
//
// # Encrypt
//
// # GenerateDataKeyPair
//
// # GenerateDataKeyPairWithoutPlaintext
//
// # GenerateDataKeyWithoutPlaintext
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [Cryptographic attestation support in KMS]: https://docs.aws.amazon.com/kms/latest/developerguide/cryptographic-attestation.html
// [Amazon Web Services Encryption SDK]: https://docs.aws.amazon.com/encryption-sdk/latest/developer-guide/
// [Amazon DynamoDB Encryption Client]: https://docs.aws.amazon.com/dynamodb-encryption-client/latest/devguide/
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [Encryption Context]: https://docs.aws.amazon.com/kms/latest/developerguide/encrypt_context.html
// [Amazon Web Services Nitro Enclaves]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/nitro-enclave.html
// [Amazon S3 client-side encryption]: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingClientSideEncryption.html
// [kms:GenerateDataKey]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
// [Amazon Web Services Nitro Enclaves SDK]: https://docs.aws.amazon.com/enclaves/latest/user/developing-applications.html#sdk
