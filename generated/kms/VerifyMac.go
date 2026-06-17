package kms

// VerifyMac is generated as a reference stub.
// Executable command wiring lives under cmd/kms.go.
//
// Verifies the hash-based message authentication code (HMAC) for a specified
// message, HMAC KMS key, and MAC algorithm. To verify the HMAC, VerifyMac
// computes an HMAC using the message, HMAC KMS key, and MAC algorithm that you
// specify, and compares the computed HMAC to the HMAC that you specify. If the
// HMACs are identical, the verification succeeds; otherwise, it fails.
// Verification indicates that the message hasn't changed since the HMAC was
// calculated, and the specified key was used to generate and verify the HMAC.
//
// HMAC KMS keys and the HMAC algorithms that KMS uses conform to industry
// standards defined in [RFC 2104].
//
// This operation is part of KMS support for HMAC KMS keys. For details, see [HMAC keys in KMS] in
// the Key Management Service Developer Guide.
//
// The KMS key that you use for this operation must be in a compatible key state.
// For details, see [Key states of KMS keys]in the Key Management Service Developer Guide.
//
// Cross-account use: Yes. To perform this operation with a KMS key in a different
// Amazon Web Services account, specify the key ARN or alias ARN in the value of
// the KeyId parameter.
//
// Required permissions: [kms:VerifyMac] (key policy)
//
// Related operations: GenerateMac
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [RFC 2104]: https://datatracker.ietf.org/doc/html/rfc2104
// [kms:VerifyMac]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
// [HMAC keys in KMS]: https://docs.aws.amazon.com/kms/latest/developerguide/hmac.html
