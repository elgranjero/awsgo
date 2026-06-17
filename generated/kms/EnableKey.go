package kms

// EnableKey is generated as a reference stub.
// Executable command wiring lives under cmd/kms.go.
//
// Sets the key state of a KMS key to enabled. This allows you to use the KMS key
// for [cryptographic operations].
//
// The KMS key that you use for this operation must be in a compatible key state.
// For details, see [Key states of KMS keys]in the Key Management Service Developer Guide.
//
// Cross-account use: No. You cannot perform this operation on a KMS key in a
// different Amazon Web Services account.
//
// Required permissions: [kms:EnableKey] (key policy)
//
// Related operations: DisableKey
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [kms:EnableKey]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [cryptographic operations]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-cryptography.html#cryptographic-operations
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
