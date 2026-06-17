package kms

// ScheduleKeyDeletion is generated as a reference stub.
// Executable command wiring lives under cmd/kms.go.
//
// Schedules the deletion of a KMS key. By default, KMS applies a waiting period
// of 30 days, but you can specify a waiting period of 7-30 days. When this
// operation is successful, the key state of the KMS key changes to PendingDeletion
// and the key can't be used in any cryptographic operations. It remains in this
// state for the duration of the waiting period. Before the waiting period ends,
// you can use CancelKeyDeletionto cancel the deletion of the KMS key. After the waiting period
// ends, KMS deletes the KMS key, its key material, and all KMS data associated
// with it, including all aliases that refer to it.
//
// Deleting a KMS key is a destructive and potentially dangerous operation. When a
// KMS key is deleted, all data that was encrypted under the KMS key is
// unrecoverable. (The only exception is a [multi-Region replica key], or an [asymmetric or HMAC KMS key with imported key material].) To prevent the use of a KMS
// key without deleting it, use DisableKey.
//
// You can schedule the deletion of a multi-Region primary key and its replica
// keys at any time. However, KMS will not delete a multi-Region primary key with
// existing replica keys. If you schedule the deletion of a primary key with
// replicas, its key state changes to PendingReplicaDeletion and it cannot be
// replicated or used in cryptographic operations. This status can continue
// indefinitely. When the last of its replicas keys is deleted (not just
// scheduled), the key state of the primary key changes to PendingDeletion and its
// waiting period ( PendingWindowInDays ) begins. For details, see [Deleting multi-Region keys] in the Key
// Management Service Developer Guide.
//
// When KMS [deletes a KMS key from an CloudHSM key store], it makes a best effort to delete the associated key material from
// the associated CloudHSM cluster. However, you might need to manually [delete the orphaned key material]from the
// cluster and its backups. [Deleting a KMS key from an external key store]has no effect on the associated external key. However,
// for both types of custom key stores, deleting a KMS key is destructive and
// irreversible. You cannot decrypt ciphertext encrypted under the KMS key by using
// only its associated external key or CloudHSM key. Also, you cannot recreate a
// KMS key in an external key store by creating a new KMS key with the same key
// material.
//
// For more information about scheduling a KMS key for deletion, see [Deleting KMS keys] in the Key
// Management Service Developer Guide.
//
// The KMS key that you use for this operation must be in a compatible key state.
// For details, see [Key states of KMS keys]in the Key Management Service Developer Guide.
//
// Cross-account use: No. You cannot perform this operation on a KMS key in a
// different Amazon Web Services account.
//
// Required permissions: kms:ScheduleKeyDeletion (key policy)
//
// # Related operations
//
// # CancelKeyDeletion
//
// # DisableKey
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [delete the orphaned key material]: https://docs.aws.amazon.com/kms/latest/developerguide/fix-keystore.html#fix-keystore-orphaned-key
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [Deleting a KMS key from an external key store]: https://docs.aws.amazon.com/kms/latest/developerguide/deleting-keys.html#delete-xks-key
// [Deleting multi-Region keys]: https://docs.aws.amazon.com/kms/latest/developerguide/deleting-keys.html#deleting-mrks
// [Deleting KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/deleting-keys.html
// [multi-Region replica key]: https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-delete.html
// [asymmetric or HMAC KMS key with imported key material]: https://docs.aws.amazon.com/kms/latest/developerguide/deleting-keys.html#import-delete-key
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
// [deletes a KMS key from an CloudHSM key store]: https://docs.aws.amazon.com/kms/latest/developerguide/deleting-keys.html#delete-cmk-keystore
