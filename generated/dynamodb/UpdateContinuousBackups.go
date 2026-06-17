package dynamodb

// UpdateContinuousBackups is generated as a reference stub.
// Executable command wiring lives under cmd/dynamodb.go.
//
// UpdateContinuousBackups enables or disables point in time recovery for the
// specified table. A successful UpdateContinuousBackups call returns the current
// ContinuousBackupsDescription . Continuous backups are ENABLED on all tables at
// table creation. If point in time recovery is enabled, PointInTimeRecoveryStatus
// will be set to ENABLED.
//
// Once continuous backups and point in time recovery are enabled, you can restore
// to any point in time within EarliestRestorableDateTime and
// LatestRestorableDateTime .
//
// LatestRestorableDateTime is typically 5 minutes before the current time. You
// can restore your table to any point in time in the last 35 days. You can set the
// RecoveryPeriodInDays to any value between 1 and 35 days.
