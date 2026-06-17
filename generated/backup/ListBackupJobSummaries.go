package backup

// ListBackupJobSummaries is generated as a reference stub.
// Executable command wiring lives under cmd/backup.go.
//
// This is a request for a summary of backup jobs created or running within the
// most recent 30 days. You can include parameters AccountID, State, ResourceType,
// MessageCategory, AggregationPeriod, MaxResults, or NextToken to filter results.
//
// This request returns a summary that contains Region, Account, State,
// ResourceType, MessageCategory, StartTime, EndTime, and Count of included jobs.
