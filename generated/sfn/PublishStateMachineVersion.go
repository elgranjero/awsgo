package sfn

// PublishStateMachineVersion is generated as a reference stub.
// Executable command wiring lives under cmd/sfn.go.
//
// Creates a [version] from the current revision of a state machine. Use versions to create
// immutable snapshots of your state machine. You can start executions from
// versions either directly or with an alias. To create an alias, use CreateStateMachineAlias.
//
// You can publish up to 1000 versions for each state machine. You must manually
// delete unused versions using the DeleteStateMachineVersionAPI action.
//
// PublishStateMachineVersion is an idempotent API. It doesn't create a duplicate
// state machine version if it already exists for the current revision. Step
// Functions bases PublishStateMachineVersion 's idempotency check on the
// stateMachineArn , name , and revisionId parameters. Requests with the same
// parameters return a successful idempotent response. If you don't specify a
// revisionId , Step Functions checks for a previously published version of the
// state machine's current revision.
//
// Related operations:
//
// # DeleteStateMachineVersion
//
// # ListStateMachineVersions
//
// [version]: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-version.html
