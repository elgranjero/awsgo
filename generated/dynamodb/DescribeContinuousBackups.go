package dynamodb

// DescribeContinuousBackups is generated as a reference stub.
// Executable command wiring lives under cmd/dynamodb.go.
//
// Checks the status of continuous backups and point in time recovery on the
// specified table. Continuous backups are ENABLED on all tables at table
// creation. If point in time recovery is enabled, PointInTimeRecoveryStatus will
// be set to ENABLED.
//
// After continuous backups and point in time recovery are enabled, you can
// restore to any point in time within EarliestRestorableDateTime and
// LatestRestorableDateTime .
//
// LatestRestorableDateTime is typically 5 minutes before the current time. You
// can restore your table to any point in time in the last 35 days. You can set the
// recovery period to any value between 1 and 35 days.
//
// You can call DescribeContinuousBackups at a maximum rate of 10 times per second.
