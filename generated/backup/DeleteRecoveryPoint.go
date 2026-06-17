package backup

// DeleteRecoveryPoint is generated as a reference stub.
// Executable command wiring lives under cmd/backup.go.
//
// Deletes the recovery point specified by a recovery point ID.
//
// If the recovery point ID belongs to a continuous backup, calling this endpoint
// deletes the existing continuous backup and stops future continuous backup.
//
// When an IAM role's permissions are insufficient to call this API, the service
// sends back an HTTP 200 response with an empty HTTP body, but the recovery point
// is not deleted. Instead, it enters an EXPIRED state.
//
// EXPIRED recovery points can be deleted with this API once the IAM role has the
// iam:CreateServiceLinkedRole action. To learn more about adding this role, see [Troubleshooting manual deletions].
//
// If the user or role is deleted or the permission within the role is removed,
// the deletion will not be successful and will enter an EXPIRED state.
//
// [Troubleshooting manual deletions]: https://docs.aws.amazon.com/aws-backup/latest/devguide/deleting-backups.html#deleting-backups-troubleshooting
