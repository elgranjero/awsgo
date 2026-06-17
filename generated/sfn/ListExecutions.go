package sfn

// ListExecutions is generated as a reference stub.
// Executable command wiring lives under cmd/sfn.go.
//
// Lists all executions of a state machine or a Map Run. You can list all
// executions related to a state machine by specifying a state machine Amazon
// Resource Name (ARN), or those related to a Map Run by specifying a Map Run ARN.
// Using this API action, you can also list all [redriven]executions.
//
// You can also provide a state machine [alias] ARN or [version] ARN to list the executions
// associated with a specific alias or version.
//
// Results are sorted by time, with the most recent execution first.
//
// If nextToken is returned, there are more results available. The value of
// nextToken is a unique pagination token for each page. Make the call again using
// the returned token to retrieve the next page. Keep all other arguments
// unchanged. Each pagination token expires after 24 hours. Using an expired
// pagination token will return an HTTP 400 InvalidToken error.
//
// This operation is eventually consistent. The results are best effort and may
// not reflect very recent updates and changes.
//
// This API action is not supported by EXPRESS state machines.
//
// [redriven]: https://docs.aws.amazon.com/step-functions/latest/dg/redrive-executions.html
// [alias]: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-alias.html
// [version]: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-version.html
