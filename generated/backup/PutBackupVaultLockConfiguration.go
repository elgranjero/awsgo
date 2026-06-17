package backup

// PutBackupVaultLockConfiguration is generated as a reference stub.
// Executable command wiring lives under cmd/backup.go.
//
// Applies Backup Vault Lock to a backup vault, preventing attempts to delete any
// recovery point stored in or created in a backup vault. Vault Lock also prevents
// attempts to update the lifecycle policy that controls the retention period of
// any recovery point currently stored in a backup vault. If specified, Vault Lock
// enforces a minimum and maximum retention period for future backup and copy jobs
// that target a backup vault.
//
// Backup Vault Lock has been assessed by Cohasset Associates for use in
// environments that are subject to SEC 17a-4, CFTC, and FINRA regulations. For
// more information about how Backup Vault Lock relates to these regulations, see
// the [Cohasset Associates Compliance Assessment.]
//
// For more information, see [Backup Vault Lock].
//
// [Cohasset Associates Compliance Assessment.]: https://docs.aws.amazon.com/aws-backup/latest/devguide/samples/cohassetreport.zip
// [Backup Vault Lock]: https://docs.aws.amazon.com/aws-backup/latest/devguide/vault-lock.html
