package kms

// GenerateMac is generated as a reference stub.
// Executable command wiring lives under cmd/kms.go.
//
// Generates a hash-based message authentication code (HMAC) for a message using
// an HMAC KMS key and a MAC algorithm that the key supports. HMAC KMS keys and the
// HMAC algorithms that KMS uses conform to industry standards defined in [RFC 2104].
//
// You can use value that GenerateMac returns in the VerifyMac operation to demonstrate
// that the original message has not changed. Also, because a secret key is used to
// create the hash, you can verify that the party that generated the hash has the
// required secret key. You can also use the raw result to implement HMAC-based
// algorithms such as key derivation functions. This operation is part of KMS
// support for HMAC KMS keys. For details, see [HMAC keys in KMS]in the Key Management Service
// Developer Guide .
//
// Best practices recommend that you limit the time during which any signing
// mechanism, including an HMAC, is effective. This deters an attack where the
// actor uses a signed message to establish validity repeatedly or long after the
// message is superseded. HMAC tags do not include a timestamp, but you can include
// a timestamp in the token or message to help you detect when its time to refresh
// the HMAC.
//
// The KMS key that you use for this operation must be in a compatible key state.
// For details, see [Key states of KMS keys]in the Key Management Service Developer Guide.
//
// Cross-account use: Yes. To perform this operation with a KMS key in a different
// Amazon Web Services account, specify the key ARN or alias ARN in the value of
// the KeyId parameter.
//
// Required permissions: [kms:GenerateMac] (key policy)
//
// Related operations: VerifyMac
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [kms:GenerateMac]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [RFC 2104]: https://datatracker.ietf.org/doc/html/rfc2104
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
// [HMAC keys in KMS]: https://docs.aws.amazon.com/kms/latest/developerguide/hmac.html
