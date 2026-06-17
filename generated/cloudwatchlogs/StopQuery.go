package cloudwatchlogs

// StopQuery is generated as a reference stub.
// Executable command wiring lives under cmd/cloudwatchlogs.go.
//
// Stops a CloudWatch Logs Insights query that is in progress. If the query has
// already ended, the operation returns an error indicating that the specified
// query is not running.
//
// This operation can be used to cancel both interactive queries and individual
// scheduled query executions. When used with scheduled queries, StopQuery cancels
// only the specific execution identified by the query ID, not the scheduled query
// configuration itself.
