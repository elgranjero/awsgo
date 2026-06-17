package sfn

// GetExecutionHistory is generated as a reference stub.
// Executable command wiring lives under cmd/sfn.go.
//
// Returns the history of the specified execution as a list of events. By default,
// the results are returned in ascending order of the timeStamp of the events. Use
// the reverseOrder parameter to get the latest events first.
//
// If nextToken is returned, there are more results available. The value of
// nextToken is a unique pagination token for each page. Make the call again using
// the returned token to retrieve the next page. Keep all other arguments
// unchanged. Each pagination token expires after 24 hours. Using an expired
// pagination token will return an HTTP 400 InvalidToken error.
//
// This API action is not supported by EXPRESS state machines.
