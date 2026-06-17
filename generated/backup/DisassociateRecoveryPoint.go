package backup

// DisassociateRecoveryPoint is generated as a reference stub.
// Executable command wiring lives under cmd/backup.go.
//
// Deletes the specified continuous backup recovery point from Backup and releases
// control of that continuous backup to the source service, such as Amazon RDS. The
// source service will continue to create and retain continuous backups using the
// lifecycle that you specified in your original backup plan.
//
// Does not support snapshot backup recovery points.
