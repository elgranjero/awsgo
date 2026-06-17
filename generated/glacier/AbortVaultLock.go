package glacier

// AbortVaultLock is generated as a reference stub.
// Executable command wiring lives under cmd/glacier.go.
//
// This operation aborts the vault locking process if the vault lock is not in the
// Locked state. If the vault lock is in the Locked state when this operation is
// requested, the operation returns an AccessDeniedException error. Aborting the
// vault locking process removes the vault lock policy from the specified vault.
//
// A vault lock is put into the InProgress state by calling InitiateVaultLock. A vault lock is put
// into the Locked state by calling CompleteVaultLock. You can get the state of a vault lock by
// calling GetVaultLock. For more information about the vault locking process, see [Amazon Glacier Vault Lock]. For more
// information about vault lock policies, see [Amazon Glacier Access Control with Vault Lock Policies].
//
// This operation is idempotent. You can successfully invoke this operation
// multiple times, if the vault lock is in the InProgress state or if there is no
// policy associated with the vault.
//
// [Amazon Glacier Vault Lock]: https://docs.aws.amazon.com/amazonglacier/latest/dev/vault-lock.html
// [Amazon Glacier Access Control with Vault Lock Policies]: https://docs.aws.amazon.com/amazonglacier/latest/dev/vault-lock-policy.html
