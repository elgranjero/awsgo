package glacier

// CompleteVaultLock is generated as a reference stub.
// Executable command wiring lives under cmd/glacier.go.
//
// This operation completes the vault locking process by transitioning the vault
// lock from the InProgress state to the Locked state, which causes the vault lock
// policy to become unchangeable. A vault lock is put into the InProgress state by
// calling InitiateVaultLock. You can obtain the state of the vault lock by calling GetVaultLock. For more
// information about the vault locking process, [Amazon Glacier Vault Lock].
//
// This operation is idempotent. This request is always successful if the vault
// lock is in the Locked state and the provided lock ID matches the lock ID
// originally used to lock the vault.
//
// If an invalid lock ID is passed in the request when the vault lock is in the
// Locked state, the operation returns an AccessDeniedException error. If an
// invalid lock ID is passed in the request when the vault lock is in the
// InProgress state, the operation throws an InvalidParameter error.
//
// [Amazon Glacier Vault Lock]: https://docs.aws.amazon.com/amazonglacier/latest/dev/vault-lock.html
