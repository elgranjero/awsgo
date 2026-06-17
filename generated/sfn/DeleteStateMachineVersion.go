package sfn

// DeleteStateMachineVersion is generated as a reference stub.
// Executable command wiring lives under cmd/sfn.go.
//
// Deletes a state machine [version]. After you delete a version, you can't call StartExecution using
// that version's ARN or use the version with a state machine [alias].
//
// Deleting a state machine version won't terminate its in-progress executions.
//
// You can't delete a state machine version currently referenced by one or more
// aliases. Before you delete a version, you must either delete the aliases or
// update them to point to another state machine version.
//
// Related operations:
//
// # PublishStateMachineVersion
//
// # ListStateMachineVersions
//
// [alias]: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-alias.html
// [version]: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-version.html
