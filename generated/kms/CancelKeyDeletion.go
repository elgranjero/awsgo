package kms

// CancelKeyDeletion is generated as a reference stub.
// Executable command wiring lives under cmd/kms.go.
//
// Cancels the deletion of a KMS key. When this operation succeeds, the key state
// of the KMS key is Disabled . To enable the KMS key, use EnableKey.
//
// For more information about scheduling and canceling deletion of a KMS key, see [Deleting KMS keys]
// in the Key Management Service Developer Guide.
//
// The KMS key that you use for this operation must be in a compatible key state.
// For details, see [Key states of KMS keys]in the Key Management Service Developer Guide.
//
// Cross-account use: No. You cannot perform this operation on a KMS key in a
// different Amazon Web Services account.
//
// Required permissions: [kms:CancelKeyDeletion] (key policy)
//
// Related operations: ScheduleKeyDeletion
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [kms:CancelKeyDeletion]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [Deleting KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/deleting-keys.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
