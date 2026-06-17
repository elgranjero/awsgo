package kms

// DisableKey is generated as a reference stub.
// Executable command wiring lives under cmd/kms.go.
//
// Sets the state of a KMS key to disabled. This change temporarily prevents use
// of the KMS key for [cryptographic operations].
//
// The KMS key that you use for this operation must be in a compatible key state.
// For more information about how key state affects the use of a KMS key, see [Key states of KMS keys]in
// the Key Management Service Developer Guide .
//
// Cross-account use: No. You cannot perform this operation on a KMS key in a
// different Amazon Web Services account.
//
// Required permissions: [kms:DisableKey] (key policy)
//
// Related operations: EnableKey
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [cryptographic operations]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-cryptography.html#cryptographic-operations
// [kms:DisableKey]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
