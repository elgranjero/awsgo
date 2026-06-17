package athena

// GetQueryRuntimeStatistics is generated as a reference stub.
// Executable command wiring lives under cmd/athena.go.
//
// Returns query execution runtime statistics related to a single execution of a
// query if you have access to the workgroup in which the query ran. Statistics
// from the Timeline section of the response object are available as soon as QueryExecutionStatus$State is
// in a SUCCEEDED or FAILED state. The remaining non-timeline statistics in the
// response (like stage-level input and output row count and data size) are updated
// asynchronously and may not be available immediately after a query completes or,
// in some cases, may not be returned. The non-timeline statistics are also not
// included when a query has row-level filters defined in Lake Formation.
