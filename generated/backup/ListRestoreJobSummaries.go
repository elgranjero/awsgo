package backup

// ListRestoreJobSummaries is generated as a reference stub.
// Executable command wiring lives under cmd/backup.go.
//
// This request obtains a summary of restore jobs created or running within the
// the most recent 30 days. You can include parameters AccountID, State,
// ResourceType, AggregationPeriod, MaxResults, or NextToken to filter results.
//
// This request returns a summary that contains Region, Account, State,
// RestourceType, MessageCategory, StartTime, EndTime, and Count of included jobs.
