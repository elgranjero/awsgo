package kms

// Decrypt is generated as a reference stub.
// Executable command wiring lives under cmd/kms.go.
//
// Decrypts ciphertext that was encrypted by a KMS key using any of the following
// operations:
//
// # Encrypt
//
// # GenerateDataKey
//
// # GenerateDataKeyPair
//
// # GenerateDataKeyWithoutPlaintext
//
// # GenerateDataKeyPairWithoutPlaintext
//
// You can use this operation to decrypt ciphertext that was encrypted under a
// symmetric encryption KMS key or an asymmetric encryption KMS key. When the KMS
// key is asymmetric, you must specify the KMS key and the encryption algorithm
// that was used to encrypt the ciphertext. For information about asymmetric KMS
// keys, see [Asymmetric KMS keys]in the Key Management Service Developer Guide.
//
// The Decrypt operation also decrypts ciphertext that was encrypted outside of
// KMS by the public key in an KMS asymmetric KMS key. However, it cannot decrypt
// symmetric ciphertext produced by other libraries, such as the [Amazon Web Services Encryption SDK]or [Amazon S3 client-side encryption]. These
// libraries return a ciphertext format that is incompatible with KMS.
//
// If the ciphertext was encrypted under a symmetric encryption KMS key, the KeyId
// parameter is optional. KMS can get this information from metadata that it adds
// to the symmetric ciphertext blob. This feature adds durability to your
// implementation by ensuring that authorized users can decrypt ciphertext decades
// after it was encrypted, even if they've lost track of the key ID. However,
// specifying the KMS key is always recommended as a best practice. When you use
// the KeyId parameter to specify a KMS key, KMS only uses the KMS key you
// specify. If the ciphertext was encrypted under a different KMS key, the Decrypt
// operation fails. This practice ensures that you use the KMS key that you intend.
//
// Whenever possible, use key policies to give users permission to call the Decrypt
// operation on a particular KMS key, instead of using IAM policies. Otherwise, you
// might create an IAM policy that gives the user Decrypt permission on all KMS
// keys. This user could decrypt ciphertext that was encrypted by KMS keys in other
// accounts if the key policy for the cross-account KMS key permits it. If you must
// use an IAM policy for Decrypt permissions, limit the user to particular KMS
// keys or particular trusted accounts. For details, see [Best practices for IAM policies]in the Key Management
// Service Developer Guide.
//
// Decrypt also supports [Amazon Web Services Nitro Enclaves] and NitroTPM, which provide attested environments in
// Amazon EC2. To call Decrypt for a Nitro enclave or NitroTPM, use the [Amazon Web Services Nitro Enclaves SDK] or any
// Amazon Web Services SDK. Use the Recipient parameter to provide the attestation
// document for the attested environment. Instead of the plaintext data, the
// response includes the plaintext data encrypted with the public key from the
// attestation document ( CiphertextForRecipient ). For information about the
// interaction between KMS and Amazon Web Services Nitro Enclaves or Amazon Web
// Services NitroTPM, see [Cryptographic attestation support in KMS]in the Key Management Service Developer Guide.
//
// The KMS key that you use for this operation must be in a compatible key state.
// For details, see [Key states of KMS keys]in the Key Management Service Developer Guide.
//
// Cross-account use: Yes. If you use the KeyId parameter to identify a KMS key in
// a different Amazon Web Services account, specify the key ARN or the alias ARN of
// the KMS key.
//
// Required permissions: [kms:Decrypt] (key policy)
//
// Related operations:
//
// # Encrypt
//
// # GenerateDataKey
//
// # GenerateDataKeyPair
//
// # ReEncrypt
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [Amazon Web Services Encryption SDK]: https://docs.aws.amazon.com/encryption-sdk/latest/developer-guide/
// [Cryptographic attestation support in KMS]: https://docs.aws.amazon.com/kms/latest/developerguide/cryptographic-attestation.html
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [kms:Decrypt]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [Asymmetric KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/symmetric-asymmetric.html
// [Amazon Web Services Nitro Enclaves]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/nitro-enclave.html
// [Amazon S3 client-side encryption]: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingClientSideEncryption.html
// [Best practices for IAM policies]: https://docs.aws.amazon.com/kms/latest/developerguide/iam-policies.html#iam-policies-best-practices
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
// [Amazon Web Services Nitro Enclaves SDK]: https://docs.aws.amazon.com/enclaves/latest/user/developing-applications.html#sdk
