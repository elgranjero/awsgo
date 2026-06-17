package sfn

// StartExecution is generated as a reference stub.
// Executable command wiring lives under cmd/sfn.go.
//
// Starts a state machine execution.
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
// - The following qualified state machine ARN refers to an alias named PROD .
//
// arn::states:::stateMachine:
//
// If you provide a qualified state machine ARN that refers to a version ARN or an
//
// alias ARN, the request starts execution for that version or alias.
//
// - The following unqualified state machine ARN refers to a state machine named
// myStateMachine .
//
// arn::states:::stateMachine:
//
// If you start an execution with an unqualified state machine ARN, Step Functions
// uses the latest revision of the state machine for the execution.
//
// To start executions of a state machine [version], call StartExecution and provide the
// version ARN or the ARN of an [alias]that points to the version.
//
// StartExecution is idempotent for STANDARD workflows. For a STANDARD workflow,
// if you call StartExecution with the same name and input as a running execution,
// the call succeeds and return the same response as the original request. If the
// execution is closed or if the input is different, it returns a 400
// ExecutionAlreadyExists error. You can reuse names after 90 days.
//
// StartExecution isn't idempotent for EXPRESS workflows.
//
// [alias]: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-alias.html
// [version]: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-version.html
