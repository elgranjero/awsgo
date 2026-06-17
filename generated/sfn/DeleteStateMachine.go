package sfn

// DeleteStateMachine is generated as a reference stub.
// Executable command wiring lives under cmd/sfn.go.
//
// Deletes a state machine. This is an asynchronous operation. It sets the state
// machine's status to DELETING and begins the deletion process. A state machine
// is deleted only when all its executions are completed. On the next state
// transition, the state machine's executions are terminated.
//
// A qualified state machine ARN can either refer to a Distributed Map state
// defined within a state machine, a version ARN, or an alias ARN.
//
// The following are some examples of qualified and unqualified state machine ARNs:
//
// - The following qualified state machine ARN refers to a Distributed Map state
// with a label mapStateLabel in a state machine named myStateMachine .
//
// arn:partition:states:region:account-id:stateMachine:myStateMachine/mapStateLabel
//
// If you provide a qualified state machine ARN that refers to a Distributed Map
//
// state, the request fails with ValidationException .
//
// - The following unqualified state machine ARN refers to a state machine named
// myStateMachine .
//
// arn:partition:states:region:account-id:stateMachine:myStateMachine
//
// This API action also deletes all [versions] and [aliases] associated with a state machine.
//
// For EXPRESS state machines, the deletion happens eventually (usually in less
// than a minute). Running executions may emit logs after DeleteStateMachine API
// is called.
//
// [aliases]: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-alias.html
// [versions]: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-version.html
