package sfn

// ListStateMachineVersions is generated as a reference stub.
// Executable command wiring lives under cmd/sfn.go.
//
// Lists [versions] for the specified state machine Amazon Resource Name (ARN).
//
// The results are sorted in descending order of the version creation time.
//
// If nextToken is returned, there are more results available. The value of
// nextToken is a unique pagination token for each page. Make the call again using
// the returned token to retrieve the next page. Keep all other arguments
// unchanged. Each pagination token expires after 24 hours. Using an expired
// pagination token will return an HTTP 400 InvalidToken error.
//
// Related operations:
//
// # PublishStateMachineVersion
//
// # DeleteStateMachineVersion
//
// [versions]: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-version.html
