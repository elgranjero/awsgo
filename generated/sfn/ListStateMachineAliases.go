package sfn

// ListStateMachineAliases is generated as a reference stub.
// Executable command wiring lives under cmd/sfn.go.
//
// Lists [aliases] for a specified state machine ARN. Results are sorted by time, with the
// most recently created aliases listed first.
//
// To list aliases that reference a state machine [version], you can specify the version
// ARN in the stateMachineArn parameter.
//
// If nextToken is returned, there are more results available. The value of
// nextToken is a unique pagination token for each page. Make the call again using
// the returned token to retrieve the next page. Keep all other arguments
// unchanged. Each pagination token expires after 24 hours. Using an expired
// pagination token will return an HTTP 400 InvalidToken error.
//
// Related operations:
//
// # CreateStateMachineAlias
//
// # DescribeStateMachineAlias
//
// # UpdateStateMachineAlias
//
// # DeleteStateMachineAlias
//
// [aliases]: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-alias.html
// [version]: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-version.html
