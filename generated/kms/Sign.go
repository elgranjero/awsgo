package kms

// Sign is generated as a reference stub.
// Executable command wiring lives under cmd/kms.go.
//
// Creates a [digital signature] for a message or message digest by using the private key in an
// asymmetric signing KMS key. To verify the signature, use the Verifyoperation, or use
// the public key in the same asymmetric KMS key outside of KMS. For information
// about asymmetric KMS keys, see [Asymmetric KMS keys]in the Key Management Service Developer Guide.
//
// Digital signatures are generated and verified by using asymmetric key pair,
// such as an RSA, ECC, or ML-DSA pair that is represented by an asymmetric KMS
// key. The key owner (or an authorized user) uses their private key to sign a
// message. Anyone with the public key can verify that the message was signed with
// that particular private key and that the message hasn't changed since it was
// signed.
//
// To use the Sign operation, provide the following information:
//
// - Use the KeyId parameter to identify an asymmetric KMS key with a KeyUsage
// value of SIGN_VERIFY . To get the KeyUsage value of a KMS key, use the DescribeKey
// operation. The caller must have kms:Sign permission on the KMS key.
//
// - Use the Message parameter to specify the message or message digest to sign.
// You can submit messages of up to 4096 bytes. To sign a larger message, generate
// a hash digest of the message, and then provide the hash digest in the Message
// parameter. To indicate whether the message is a full message, a digest, or an
// ML-DSA EXTERNAL_MU, use the MessageType parameter.
//
// - Choose a signing algorithm that is compatible with the KMS key.
//
// When signing a message, be sure to record the KMS key and the signing
// algorithm. This information is required to verify the signature.
//
// Best practices recommend that you limit the time during which any signature is
// effective. This deters an attack where the actor uses a signed message to
// establish validity repeatedly or long after the message is superseded.
// Signatures do not include a timestamp, but you can include a timestamp in the
// signed message to help you detect when its time to refresh the signature.
//
// To verify the signature that this operation generates, use the Verify operation. Or
// use the GetPublicKeyoperation to download the public key and then use the public key to
// verify the signature outside of KMS.
//
// The KMS key that you use for this operation must be in a compatible key state.
// For details, see [Key states of KMS keys]in the Key Management Service Developer Guide.
//
// Cross-account use: Yes. To perform this operation with a KMS key in a different
// Amazon Web Services account, specify the key ARN or alias ARN in the value of
// the KeyId parameter.
//
// Required permissions: [kms:Sign] (key policy)
//
// Related operations: Verify
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [digital signature]: https://en.wikipedia.org/wiki/Digital_signature
// [Asymmetric KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/symmetric-asymmetric.html
// [kms:Sign]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
