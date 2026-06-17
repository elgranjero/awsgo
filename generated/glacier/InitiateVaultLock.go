package glacier

// InitiateVaultLock is generated as a reference stub.
// Executable command wiring lives under cmd/glacier.go.
//
// This operation initiates the vault locking process by doing the following:
//
// - Installing a vault lock policy on the specified vault.
//
// - Setting the lock state of vault lock to InProgress .
//
// - Returning a lock ID, which is used to complete the vault locking process.
//
// You can set one vault lock policy for each vault and this policy can be up to
// 20 KB in size. For more information about vault lock policies, see [Amazon Glacier Access Control with Vault Lock Policies].
//
// You must complete the vault locking process within 24 hours after the vault
// lock enters the InProgress state. After the 24 hour window ends, the lock ID
// expires, the vault automatically exits the InProgress state, and the vault lock
// policy is removed from the vault. You call CompleteVaultLockto complete the vault locking
// process by setting the state of the vault lock to Locked .
//
// After a vault lock is in the Locked state, you cannot initiate a new vault lock
// for the vault.
//
// You can abort the vault locking process by calling AbortVaultLock. You can get the state of
// the vault lock by calling GetVaultLock. For more information about the vault locking
// process, [Amazon Glacier Vault Lock].
//
// If this operation is called when the vault lock is in the InProgress state, the
// operation returns an AccessDeniedException error. When the vault lock is in the
// InProgress state you must call AbortVaultLock before you can initiate a new vault lock
// policy.
//
// [Amazon Glacier Vault Lock]: https://docs.aws.amazon.com/amazonglacier/latest/dev/vault-lock.html
// [Amazon Glacier Access Control with Vault Lock Policies]: https://docs.aws.amazon.com/amazonglacier/latest/dev/vault-lock-policy.html
