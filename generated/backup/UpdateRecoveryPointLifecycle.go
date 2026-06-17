package backup

// UpdateRecoveryPointLifecycle is generated as a reference stub.
// Executable command wiring lives under cmd/backup.go.
//
// Sets the transition lifecycle of a recovery point.
//
// The lifecycle defines when a protected resource is transitioned to cold storage
// and when it expires. Backup transitions and expires backups automatically
// according to the lifecycle that you define.
//
// Resource types that can transition to cold storage are listed in the [Feature availability by resource] table.
// Backup ignores this expression for other resource types.
//
// Backups transitioned to cold storage must be stored in cold storage for a
// minimum of 90 days. Therefore, the “retention” setting must be 90 days greater
// than the “transition to cold after days” setting. The “transition to cold after
// days” setting cannot be changed after a backup has been transitioned to cold.
//
// If your lifecycle currently uses the parameters DeleteAfterDays and
// MoveToColdStorageAfterDays , include these parameters and their values when you
// call this operation. Not including them may result in your plan updating with
// null values.
//
// This operation does not support continuous backups.
//
// [Feature availability by resource]: https://docs.aws.amazon.com/aws-backup/latest/devguide/backup-feature-availability.html#features-by-resource
