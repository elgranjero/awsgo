package sfn

// UpdateStateMachineAlias is generated as a reference stub.
// Executable command wiring lives under cmd/sfn.go.
//
// Updates the configuration of an existing state machine [alias] by modifying its
// description or routingConfiguration .
//
// You must specify at least one of the description or routingConfiguration
// parameters to update a state machine alias.
//
// UpdateStateMachineAlias is an idempotent API. Step Functions bases the
// idempotency check on the stateMachineAliasArn , description , and
// routingConfiguration parameters. Requests with the same parameters return an
// idempotent response.
//
// This operation is eventually consistent. All StartExecution requests made within a few
// seconds use the latest alias configuration. Executions started immediately after
// calling UpdateStateMachineAlias may use the previous routing configuration.
//
// Related operations:
//
// # CreateStateMachineAlias
//
// # DescribeStateMachineAlias
//
// # ListStateMachineAliases
//
// # DeleteStateMachineAlias
//
// [alias]: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-alias.html
