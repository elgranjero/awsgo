package sfn

// CreateStateMachineAlias is generated as a reference stub.
// Executable command wiring lives under cmd/sfn.go.
//
// Creates an [alias] for a state machine that points to one or two [versions] of the same state
// machine. You can set your application to call StartExecutionwith an alias and update the
// version the alias uses without changing the client's code.
//
// You can also map an alias to split StartExecution requests between two versions of a state
// machine. To do this, add a second RoutingConfig object in the
// routingConfiguration parameter. You must also specify the percentage of
// execution run requests each version should receive in both RoutingConfig
// objects. Step Functions randomly chooses which version runs a given execution
// based on the percentage you specify.
//
// To create an alias that points to a single version, specify a single
// RoutingConfig object with a weight set to 100.
//
// You can create up to 100 aliases for each state machine. You must delete unused
// aliases using the DeleteStateMachineAliasAPI action.
//
// CreateStateMachineAlias is an idempotent API. Step Functions bases the
// idempotency check on the stateMachineArn , description , name , and
// routingConfiguration parameters. Requests that contain the same values for these
// parameters return a successful idempotent response without creating a duplicate
// resource.
//
// Related operations:
//
// # DescribeStateMachineAlias
//
// # ListStateMachineAliases
//
// # UpdateStateMachineAlias
//
// # DeleteStateMachineAlias
//
// [versions]: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-version.html
// [alias]: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-alias.html
